package comment

import (
	"gorm.io/gorm"
)

type Service interface {
	GetComment(ID uint) (Comment, error)
	GetCommentsBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

type service struct {
	DB *gorm.DB
}

func (s *service) GetComment(ID uint) (Comment, error) {
	var comment Comment

	if result := s.DB.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

func (s *service) GetCommentsBySlug(slug string) ([]Comment, error) {
	var comments []Comment

	if result := s.DB.Where("slug = ?", slug).Find(&comments); result.Error != nil {
		return []Comment{}, result.Error
	}

	return comments, nil
}

func (s *service) PostComment(comment Comment) (Comment, error) {
	if result := s.DB.Save(&comment); result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

func (s *service) UpdateComment(ID uint, newComment Comment) (Comment, error) {
	comment, err := s.GetComment(ID)

	if err != nil {
		return Comment{}, err
	}

	if result := s.DB.Model(&comment).Updates(newComment); result.Error != nil {
		return Comment{}, err
	}

	return newComment, nil
}

func (s *service) DeleteComment(ID uint) error {
	if result := s.DB.Delete(&Comment{}, ID); result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *service) GetAllComments() ([]Comment, error) {
	var comments []Comment

	if result := s.DB.Find(&comments); result.Error != nil {
		return []Comment{}, result.Error
	}

	return comments, nil
}

func NewService(db *gorm.DB) Service {
	return &service{
		DB: db,
	}
}
