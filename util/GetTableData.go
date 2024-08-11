package util

import (
	"bytes"
	"encoding/binary"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"os"
	"path/filepath"
	"reflect"
)

func GetTableData[T proto.Message](tablePath string, tableData T) error {
	tableType := reflect.TypeOf(tableData)

	fileName := tableType.Elem().Name() + ".bytes"
	fileData, err := os.ReadFile(filepath.Join(tablePath, fileName))
	if err != nil {
		return errors.WithStack(err)
	}

	var head uint32
	err = binary.Read(bytes.NewReader(fileData[:4]), binary.LittleEndian, &head)
	if err != nil {
		return errors.WithStack(err)
	}

	err = proto.Unmarshal(fileData[head+4:], tableData)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
