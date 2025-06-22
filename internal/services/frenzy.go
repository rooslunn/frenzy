package services

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Frenzy struct {
	Original        string
	OriginalStory   string
	Translated      string
	TranslatedStory string
	ImagePath       string
	Langpair        string
}

const (
	frenzySep   = ";"
	ErrEmptyRaw = "empty raw frenzy"
	ImageDir    = "data/img"
	DataFile = "data/frenzy"
)

var tgEscapeRunes = []rune{'.', '-', '!'}

func dayImage() (string, error) {
	
	img_count, err := filesCountByMask(filepath.Join(ImageDir, "*.png")) 
	if err != nil {
		return "", err
	}

	dayOfYear := time.Now().YearDay()
	index := (dayOfYear - 1) % img_count 

	fn := fmt.Sprintf("%d.png", index)
	fPath := filepath.Join(ImageDir, fn)

	if !FileExists(fPath) {
		return "", fmt.Errorf("%s: %w", fPath, ErrFileMissing)
	}

	return fPath, nil
}

func readFrenzyRotate(fn string) (string, error) {
	file, err := os.Open(fn)
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	if len(lines) == 0 {
		return "", fmt.Errorf("file is empty")
	}

	dayOfYear := time.Now().YearDay()

	index := (dayOfYear - 1) % len(lines)

	return lines[index], nil
}

func FetchFrenzy(ctx context.Context, aiKey string) (Frenzy, error) {

	// frenzyUncooked := "Compassion fatique; Усталость от сочувствия"
	frenzyUncooked, err := readFrenzyRotate(DataFile)
	if err != nil {
		return Frenzy{}, err
	}

	frenzy, err := frenzyFromRaw(ctx, aiKey, frenzyUncooked)
	if err != nil {
		return frenzy, err
	}

	return frenzy, nil
}

func bimboFrenzy(frenzy Frenzy) string {
	originalStory := frenzyBold(frenzy.OriginalStory, frenzy.Original)
	translatedStory := frenzyBold(frenzy.TranslatedStory, frenzy.Translated)

	bimbo := fmt.Sprintf("*%s*\n_%s_\n---\n%s---\n%s", frenzy.Original, frenzy.Translated, originalStory, translatedStory)
	bimbo = EscapeMultipleChars(bimbo, tgEscapeRunes)

	return bimbo
}

func frenzyFromRaw(ctx context.Context, aiKey, raw string) (Frenzy, error) {

	parts := strings.Split(raw, frenzySep)
	l := len(parts)

	frenzy := Frenzy{}

	if l == 0 {
		return frenzy, errors.New(ErrEmptyRaw)
	}

	frenzy.Original = cleanString(parts[0])
	story, err := GenerateFenzyText(ctx, aiKey, "en", frenzy.Original)
	if err != nil {
		return frenzy, err
	}
	frenzy.OriginalStory = story

	if l == 2 {
		frenzy.Translated = cleanString(parts[1])
		story, err = GenerateFenzyText(ctx, aiKey, "ru", frenzy.Translated)
		if err != nil {
			return frenzy, err
		}
		frenzy.TranslatedStory = story
	}

	image, err := dayImage()
	if err != nil {
		return frenzy, err
	}

	frenzy.ImagePath = image 

	return frenzy, nil
}
