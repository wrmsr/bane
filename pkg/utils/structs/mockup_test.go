/*
TODO:
 - om
  - field
   - default: ta.Union[ta.Any, MISSING_TYPE] = MISSING,
   - default_factory: ta.Union[ta.Callable[[], ta.Any], MISSING_TYPE] = MISSING,
   - init: bool = True,
   - repr: bool = True,
   - hash: ta.Optional[bool] = None,
   - compare: bool = True,
   - metadata: ta.Optional[ta.Mapping[ta.Any, ta.Any]] = None,
   - new: ta.Union[ta.Callable[[], ta.Any], MISSING_TYPE] = MISSING,
   - doc: ta.Optional[str] = None,
   - mangled: ta.Optional[str] = None,
   - size: ta.Optional[ta.Any] = None,
   - frozen: ta.Optional[bool] = None,
   - kwonly: bool = False,
   - coerce: ta.Optional[ta.Union[bool, ta.Callable[[ta.Any], ta.Any]]] = None,
   - derive: ta.Optional[ta.Callable[..., ta.Any]] = None,
   - check: ta.Optional[ta.Union[bool, ta.Callable[[ta.Any], bool]]] = None,
   - check_type: ta.Union[ta.Type, ta.Tuple, None] = None,
   - validate: ta.Optional[ta.Union[bool, ta.Callable[[ta.Any], None]]] = None,
   - repr_if: ta.Optional[ta.Callable[[ta.Any], bool]] = None,
   - repr_fn: ta.Optional[ta.Callable[[ta.Any], str]] = None,
  - class
   - init: ta.Union[bool, MISSING_TYPE] = MISSING,  # True
   - repr: ta.Union[bool, MISSING_TYPE] = MISSING,  # True
   - eq: ta.Union[bool, MISSING_TYPE] = MISSING,  # True
   - order: ta.Union[bool, MISSING_TYPE] = MISSING,  # False
   - unsafe_hash: ta.Union[bool, MISSING_TYPE] = MISSING,  # False
   - frozen: ta.Union[bool, MISSING_TYPE] = MISSING,  # False
   - metadata: ta.Optional[ta.Mapping[ta.Any, ta.Any]] = None,
   - validate: ta.Union[bool, MISSING_TYPE] = MISSING,
   - field_attrs: ta.Union[bool, MISSING_TYPE] = MISSING,
   - kwonly: ta.Union[None, bool, MISSING_TYPE] = MISSING,
   - cache_hash: ta.Union[bool, str, MISSING_TYPE] = MISSING,
   - strict_eq: ta.Union[bool, MISSING_TYPE] = MISSING,
   - pickle: ta.Union[bool, MISSING_TYPE] = MISSING,
   - reorder: ta.Union[bool, MISSING_TYPE] = MISSING,
   - allow_setattr: ta.Union[bool, str, ta.Iterable[str], MISSING_TYPE] = MISSING,
   - iterable: ta.Union[bool, MISSING_TYPE] = MISSING,
   - mangler: ta.Union[Mangler, MISSING_TYPE] = MISSING,
   - aspects: ta.Union[None, ta.Sequence[ta.Any], MISSING_TYPE] = MISSING,
   - confer: ta.Union[None, ta.Sequence[str], ta.Mapping[str, ta.Any], MISSING_TYPE] = MISSING,
  - build
  - replace
  - coerce
  - validate
  - freeze?
  - dispatch
   - match?
 - attrs
  -
 - mapstruct
  -
 - json accelerator? StructTool accelerator?
*/
package structs

import "testing"

//

type RootDef interface {
	isRootDef()
}

func Def(defs ...RootDef) any {
	return nil
}

//

type StructOpt interface {
	isStructOpt()
}

type StructDef struct {
	Name string
	Opts []StructOpt
}

func (sd StructDef) isRootDef() {}

func Struct(name string, opts ...StructOpt) StructDef {
	return StructDef{
		Name: name,
		Opts: opts,
	}
}

//

type FieldOpt interface {
	isFieldOpt()
}

type FieldDef struct {
	Name string
	Opts []FieldOpt
}

func (fd FieldDef) isStructOpt() {}

func Field(name string, opts ...FieldOpt) FieldDef {
	return FieldDef{
		Name: name,
		Opts: opts,
	}
}

//

var _ = Def(
	Struct("Foo",
		Field("bar")))

//

func TestMockup(t *testing.T) {

}
