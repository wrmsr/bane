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
	type_
}

//

type Struct struct {
	value
}
