package graph

import (
	"encoding/base64"
	"fmt"
	"mybanks-api/ent"
	"strconv"
)

func encodeCursorFromInt(id int) ent.Cursor {
	return ent.Cursor{ID: id}
}
func decodeCursorToInt(cursor string) (int, error) {
	bytes, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return 0, fmt.Errorf("invalid cursor: %w", err)
	}
	return strconv.Atoi(string(bytes))
}

// EncodeCursor сериализует любую структуру курсора в base64-строку
func EncodeCursor(id int) (string, error) {
	// Преобразуем ID в байтовый срез
	cursorBytes := []byte(fmt.Sprintf("%d", id))

	// Кодируем в Base64
	cursor := base64.StdEncoding.EncodeToString(cursorBytes)

	return cursor, nil
}

// DecodeCursor декодирует курсор из Base64 в исходный ID
func DecodeCursor(cursor string) (int, error) {
	// Декодируем строку Base64 в байтовый срез
	decodedBytes, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return 0, fmt.Errorf("failed to decode cursor: %v", err)
	}

	// Преобразуем байтовый срез в строку и затем в целочисленный ID
	id, err := strconv.Atoi(string(decodedBytes))
	if err != nil {
		return 0, fmt.Errorf("failed to convert decoded cursor to int: %v", err)
	}

	return id, nil
}
