package utils

import (
	"fmt"
	"github.com/khalilsayhi/treasure-hunt/domain/types"
	"log"
	"strings"
)

func WriteSimulationResultToFile(boardCells [][]types.Cell, width, height int, initialMapPath string) {
	filenameWithoutExt, _ := strings.CutSuffix(initialMapPath, ".txt")

	// En lisant L'ennoncé j'ai compris qu'il faut afficher
	// les elements regroupés et pas dans l'ordre de leur apparition avec des lignes de commentaires au debut de chaque groupe
	// sinon il suffit juste d'utiliser outputLines pour toute les lignes
	outputLines := []string{
		"# {C comme Carte} - {Nb. de case en largeur} - {Nb. de case en hauteur}",
		"C - " + fmt.Sprint(width) + " - " + fmt.Sprint(height),
	}
	mountainLines := []string{
		"# {M comme Montagne} - {Axe horizontal} - {Axe vertical}",
	}
	treasureLines := []string{
		"# {T comme Trésor} - {Axe horizontal} - {Axe vertical} - {Nb. de trésors restants}",
	}

	adventurerLines := []string{
		"# {A comme Aventurier} - {Nom de l’aventurier} - {Axe horizontal} - {Axevertical} - {Orientation} - {Nb. trésors ramassés}",
	}

	for i := range boardCells {
		for j, content := range boardCells[i] {
			if content.Adventurer != nil {
				adventurerLines = append(adventurerLines, fmt.Sprintf("A - %s - %d - %d - %s - %d", content.Adventurer.Name, j, i, content.Adventurer.Orientation, content.Adventurer.TreasureCount))
			}
			if content.Treasure != nil {
				treasureLines = append(treasureLines, fmt.Sprintf("T - %d - %d - %d", j, i, content.Treasure.Amount))
			} else if content.Mountain != nil {
				mountainLines = append(mountainLines, fmt.Sprintf("M - %d - %d", j, i))
			}
		}
	}

	if len(mountainLines) > 1 {
		outputLines = append(outputLines, mountainLines...)
	}
	if len(treasureLines) > 1 {
		outputLines = append(outputLines, treasureLines...)
	}
	if len(adventurerLines) > 1 {
		outputLines = append(outputLines, adventurerLines...)
	}

	resFileName := fmt.Sprintf("%s_result.txt", filenameWithoutExt)
	if err := WriteLineListToFile(resFileName, outputLines); err != nil {
		log.Printf("Error writing to file: %+v\n", err)
	} else {
		log.Println("Game state written to simulation_result.txt")
	}
}
