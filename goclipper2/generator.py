import warnings
import os
from pycparser import c_ast, parse_file
from typing import Tuple, List
from dataclasses import dataclass
from predef import methods, header

# types_mapping define how each type of C program is treated inside golang templating
types_mapping = {
    'int64_t': 'int64',
    'double': 'float64',
    'int': 'int64',
    'int64': 'int64',
    'size_t': 'int64',
    'ClipperFillRule': 'ClipperFillRule',
    'ClipperClipType': 'ClipperClipType',
    'CLipperPathType': 'ClipperPathType',
    'ClipperJoinType': 'ClipperJoinType',
    'ClipperEndType': 'ClipperEndType',
    'ClipperPointInPolygonResult': 'ClipperPointInPolygonResult'
}


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


def is_complex_type_name(name: str):
    try:
        types_mapping[name]

        return False
    except KeyError:
        return True


def update_name_sign(param: Type):
    """
    updates param type_name for function signature
    """

    try:
        return types_mapping[param.type_name]
    except KeyError:
        return param.type_name


def update_name_pass(param: Type):
    """
    update_name returns the same name if it is not compound type, otherwise adds call to inner C struct
    """

    try:
        types_mapping[param.type_name]

        return f"C.{param.type_name}({param.name})"
    except KeyError:
        return f"{param.name}.P"


def is_constructor(name: str) -> bool:
    """
    is_constructor checks whether function name matches pattern for constructor function
    """
    return len(name.split("_")) == 2


def template_constructor(functype: Type, params: List[Type], has_mem: bool):
    param_signature = ", ".join([
        f"{p.name} {update_name_sign(p)}" for p in params
    ])

    param_call = ", ".join(
        [("*" if p.is_ptr and not is_complex_type_name(p.type_name) else "") + update_name_pass(p) for p in params])

    return f"""
    func {functype.name.capitalize()}({param_signature}) {"*" if functype.is_ptr else ""}{functype.type_name} {{
        {"var mem unsafe.Pointer = C.malloc(0)" if has_mem else ""}

        return &{functype.type_name}{{
            P: C.{functype.name}({"mem, " if has_mem else ""}{param_call}),
        }}
    }}
    """


def is_method(func: Type, params: List[Type]):
    """
    declaration of method: 1st non-mem argument has the same name as function name
    clipper_pathd_rect_clip_line(void *mem, ClipperRectD)
    """

    fname = func.name.replace("_", "")
    return len(params) > 0 and fname.lower().startswith(params[0].type_name.lower())


def template_method(functype: Type, params: List[Type], has_mem: bool):
    # first param is said to be struct receiver
    receiver, params = params[0], params[1:]

    param_signature = ", ".join([
        f"{p.name} {update_name_sign(p)}" for p in params
    ])

    param_call = ", ".join([f"{receiver.name}.P"] +
                           [update_name_pass(p) for p in params])

    is_complex_return_type = is_complex_type_name(functype.type_name)

    if functype.type_name == "void":
        ret_templ = f"""C.{functype.name}({"mem, " if has_mem else ""}{param_call})"""
    elif is_complex_return_type:
        ret_templ = f"""
        return {"&" if functype.is_ptr else ""}{functype.type_name}{{
            P: C.{functype.name}({"mem, " if has_mem else ""}{param_call}),
        }}
        """
    else:
        ret_templ = f"""return {update_name_sign(functype)}(C.{functype.name}({"mem, " if has_mem else ""}{param_call}))"""

    return f"""
    func ({receiver.name} *{receiver.type_name}){functype.name.capitalize()}({param_signature}) {"*" if functype.is_ptr else ""}{update_name_sign(functype) if functype.type_name != "void" else ""} {{
        {"var mem unsafe.Pointer = C.malloc(0)" if has_mem else ""}

        {ret_templ}
    }}
    """


class FuncDefVisitor(c_ast.NodeVisitor):
    def __init__(self) -> None:
        super().__init__()

        self.constructor_writer = open('constructor.go', "w")
        self.methods_writer = open('methods.go', 'w')

        self.constructor_writer.write(header)
        self.methods_writer.write(header)

    def visit_FuncDecl(self, node):
        if type(node.type) == c_ast.PtrDecl and type(node.type.type) == c_ast.PtrDecl:
            warnings.warn(
                f"manually add function: {node.type.type.type.declname}")
            return

        function = parse_decl(node)

        # try and check function in predefined
        try:
            self.methods_writer.write(methods[function.name])
            return
        except KeyError:
            ...

        if not node.args:
            return

        params = []
        has_mem = False
        any_broken = False

        for param in node.args:
            if param.name == 'mem':
                has_mem = True
                continue

            parameter = parse_decl(param)

            if parameter == None:
                any_broken = True

            params.append(parameter)

        # we have parsed function signature, now decide it is constructor or method and template

        if any_broken:
            warnings.warn(f"broken function: {function}")
            return

        if is_constructor(function.name):
            self.constructor_writer.write(
                template_constructor(function, params, has_mem))
        elif is_method(function, params):
            self.methods_writer.write(
                template_method(function, params, has_mem))

    def close(self):
        self.constructor_writer.close()
        self.methods_writer.close()


def show_func_defs(filename):
    # Note that cpp is used. Provide a path to your own cpp or
    # make sure one exists in PATH.
    ast = parse_file(filename,
                     use_cpp=True,
                     cpp_args=r'-I../pycparser/utils/fake_libc_include')

    v = FuncDefVisitor()
    v.visit(ast)

    v.close()


if __name__ == "__main__":
    filename = '/home/leo/spycer/goclipper2/vendor/clipper2c/include/clipper2c/clipper2c.h'

    show_func_defs(filename)
