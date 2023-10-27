package noteHandler

import (
	"fmt"
	"time"

	"github.com/admzulfikar97/learn_gofiber_app_server/database"
	"github.com/admzulfikar97/learn_gofiber_app_server/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []model.Note
	var notesResponse []model.Note

	db.Find(&notes)

	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
	}

	for _, note := range notes {
		noteTemp := note
		// noteTemp.UpdateAt = note.UpdatedAt.Format("2006-01-02T15:04:05Z")
		notesResponse = append(notesResponse, noteTemp)
	}

	return c.Status(200).Status(200).JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": notesResponse})
}

func CreateNotes(c *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)
	timeNow := time.Now().Format("2023-10-26T16:37:39Z")

	var countNotes []model.Note
	db.Find(&countNotes)

	// Store the body in the note and return error if encountered
	err := c.BodyParser(note)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// note.ID = len(countNotes) + 1
	fmt.Println(note.ID)
	// Add a uuid to the note
	note.UUID = uuid.New()
	note.UpdatedAt = timeNow
	note.CreatedAt = timeNow
	// Create the Note and return error if encountered
	err = db.Create(&note).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
	}

	// note.UpdateAt = note.UpdatedAt.Format("2006-01-02T15:04:05Z")

	// Return the created note
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Created Note", "data": note})
}

func GetNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	// Read the param noteId
	id := c.Params("noteID")

	// Find the note with the given Id
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	if note.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	// note.UpdateAt = note.UpdatedAt.Format("2006-01-02T15:04:05Z")

	// Return the note with the Id
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": note})
}

func UpdateNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	te := c.Request()
	fmt.Println(string(te.Body()))

	// Read the param noteId
	id := c.Params("noteID")

	// Find the note with the given Id
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	if note.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	// Store the body containing the updated data and return error if encountered
	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// if updateNoteData.UpdatedAt.Format("2006-01-02T15:04:05Z") != note.UpdatedAt.Format("2006-01-02T15:04:05Z") {
	// 	return c.Status(404).JSON(fiber.Map{"status": "error", "message": "note has updated by other", "data": nil})
	// }

	// Edit the note
	note.Title = updateNoteData.Title
	note.SubTitle = updateNoteData.SubTitle
	note.Text = updateNoteData.Text
	// note.UpdatedAt = time.Now()

	// Save the Changes
	db.Save(&note)

	// note.UpdateAt = note.UpdatedAt.Format("2006-01-02T15:04:05Z")

	// Return the updated note
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": note})
}

func DeleteNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	// Read the param noteId
	id := c.Params("noteID")

	// Find the note with the given Id
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	if note.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	// Store the body containing the updated data and return error if encountered
	var deleteNoteData updateNote
	err := c.BodyParser(&deleteNoteData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// if deleteNoteData.UpdatedAt.Format("2006-01-02T15:04:05Z") != note.UpdatedAt.Format("2006-01-02T15:04:05Z") {
	// 	return c.Status(404).JSON(fiber.Map{"status": "error", "message": "note has updated by other", "data": nil})
	// }

	// Delete the note and return error if encountered
	err = db.Delete(&note, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete note", "data": nil})
	}

	// Return success message
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Deleted Note"})
}

type updateNote struct {
	Title     string    `json:"title"`
	SubTitle  string    `json:"sub_title"`
	Text      string    `json:"Text"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
