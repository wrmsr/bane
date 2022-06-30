package tower

import "reflect"

//

type FieldType struct {
	s *StructType
	f reflect.StructField
}

//

type Field struct {
	s *Struct
	f reflect.StructField
}

//

type StructType struct {
	BaseType
}

//

type Struct struct {
	BaseValue
}
