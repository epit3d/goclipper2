from pycparser import c_ast, parse_file
from typing import Tuple, List
from dataclasses import dataclass


@dataclass
class Type:
    name: str
    type_name: str
    is_ptr: bool
    is_struct: bool


def parse_decl(node: c_ast.Decl) -> Type:
    """
    parse_decl returns (name, type, used_pointer, is_struct) of some declaration
    might be declaration of a function or parameter
    """
    # check if it is ** and skip
    if type(node.type) == c_ast.PtrDecl and type(node.type.type) == c_ast.PtrDecl:
        return None

    is_ptr = False

    if type(node.type) == c_ast.PtrDecl:
        node = node.type
        is_ptr = True

    name = node.type.declname

    is_struct = False
    try:
        name_type = node.type.type.names[0]
    except:
        name_type = node.type.type.name
        is_struct = True

    return Type(name, name_type, is_ptr, is_struct)


def template(struct_name, func_name, parameters, return_type):
    return f"""
    func (clipperObj {struct_name}) {func_name}({parameters}) {return_type} {{
        
        return {return_type} // not actually, rather object
    }}
    """


def update_name_sign(param: Type):
    simple = {
        'int64_t': 'int64',
        'double': 'float64',
        'int': 'int64',
        'int64': 'int64'
    }

    try:
        return simple[param.type_name]
    except KeyError:
        return param.type_name


def update_name_pass(param: Type):
    """
    update_name returns the same name if it is not compound type, otherwise adds call to inner C struct
    """

    simple = {
        'int64_t': 'int64',
        'double': 'float64',
        'int': 'int64',
        'int64': 'int64',
    }

    try:
        simple[param.type_name]

        return f"C.{param.type_name}({param.name})"
    except KeyError:
        return f"{param.name}.P"


def is_constructor(name: str) -> bool:
    """
    is_constructor checks whether function name matches pattern for constructor function
    """
    return len(name.split("_")) == 2


def template_constructor(functype: Type, params: List[Type]):
    param_signature = ", ".join([
        f"{p.name} {update_name_sign(p)}" for p in params
    ])

    param_call = ", ".join([update_name_pass(p) for p in params])

    return f"""
    func New_{functype.name}({param_signature}) *{functype.type_name} {{
        var mem unsafe.Pointer = C.malloc(0)

        return &{functype.type_name}{{
            P: C.{functype.name}(mem, {param_call}),
        }}
    }}
    """


class FuncDefVisitor(c_ast.NodeVisitor):
    def visit_FuncDecl(self, node):
        # print(node)
        # print(node.type)

        if type(node.type) == c_ast.PtrDecl and type(node.type.type) == c_ast.PtrDecl:
            # print(f"manually add function {node.type.type.type.declname}")
            return

        function = parse_decl(node)
        if not is_constructor(function.name):
            # skip not constructors
            return

        # print(function)
        args = node.args

        if not args:
            return

        params = []

        for param in args:
            if param.name == 'mem':
                continue

            parameter = parse_decl(param)

            # print('\t', parameter)

            params.append(parameter)

        # print(args)

        print(template_constructor(function, params))

        # exit(1)


def show_func_defs(filename):
    # Note that cpp is used. Provide a path to your own cpp or
    # make sure one exists in PATH.
    ast = parse_file(filename,
                     use_cpp=True,
                     cpp_args=r'-I../pycparser/utils/fake_libc_include')

    v = FuncDefVisitor()
    v.visit(ast)


if __name__ == "__main__":
    filename = '/home/leo/spycer/goclipper2/vendor/clipper2c/include/clipper2c/clipper2c.h'

    show_func_defs(filename)
