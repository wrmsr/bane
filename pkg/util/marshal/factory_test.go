package marshal

import "testing"

func TestDefaultFactory(t *testing.T) {
	mf := NewCompositeFactory(
		FirstComposite,
		NewPrimitiveMarshalerFactory(),
		NewPointerMarshalerFactory(),
		NewIndexMarshalerFactory(),
		NewMapMarshalerFactory(),
		NewBase64MarshalerFactory(),
		NewOptionalMarshalerFactory(),
	)

	uf := NewCompositeFactory(
		FirstComposite,
		NewConvertPrimitiveUnmarshalerFactory(),
		NewPointerUnmarshalerFactory(),
		NewIndexUnmarshalerFactory(),
		NewMapUnmarshalerFactory(),
		NewBase64UnmarshalerFactory(),
		NewOptionalUnmarshalerFactory(),
	)
}
