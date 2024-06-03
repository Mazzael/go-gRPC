package service

import (
	"context"

	database "github.com/Mazzael/gRPC/internal"
	"github.com/Mazzael/gRPC/internal/pb"
)

type CourseService struct {
	pb.UnimplementedCourseServiceServer
	CourseDB database.Course
}

func NewCourseService(courseDB database.Course) *CourseService {
	return &CourseService{
		CourseDB: courseDB,
	}
}

func (c *CourseService) CreateCourse(ctx context.Context, in *pb.CreateCourseRequest) (*pb.Course, error) {
	course, err := c.CourseDB.Create(in.Name, in.Description, in.CategoryId)

	if err != nil {
		return nil, err
	}

	courseResponse := &pb.Course{
		Id:          course.ID,
		Name:        course.Name,
		Description: course.Description,
	}

	return courseResponse, nil
}

func (c *CourseService) ListCourses(ctx context.Context, in *pb.Blank) (*pb.CourseList, error) {
	courses, err := c.CourseDB.FindAll()

	if err != nil {
		return nil, err
	}

	var coursesResponse []*pb.Course

	for _, course := range courses {
		coursesResponse = append(coursesResponse, &pb.Course{
			Id:          course.ID,
			Name:        course.Name,
			Description: course.Description,
		})
	}

	return &pb.CourseList{
		Courses: coursesResponse,
	}, nil
}

func (c *CourseService) GetCourse(ctx context.Context, in *pb.GetCourseRequest) (*pb.Course, error) {
	course, err := c.CourseDB.FindById(in.Id)

	if err != nil {
		return nil, err
	}

	return &pb.Course{
		Id:          course.ID,
		Name:        course.Name,
		Description: course.Description,
	}, nil
}

func (c *CourseService) GetCoursesByCategoryId(ctx context.Context, in *pb.GetCoursesByCategoryIdRequest) (*pb.CourseList, error) {
	courses, err := c.CourseDB.FindByCategoryId(in.CategoryId)

	if err != nil {
		return nil, err
	}

	var coursesResponse []*pb.Course

	for _, course := range courses {
		coursesResponse = append(coursesResponse, &pb.Course{
			Id:          course.ID,
			Name:        course.Name,
			Description: course.Description,
		})
	}

	return &pb.CourseList{
		Courses: coursesResponse,
	}, nil
}
