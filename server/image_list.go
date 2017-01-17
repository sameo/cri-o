package server

import (
	"github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	pb "k8s.io/kubernetes/pkg/kubelet/api/v1alpha1/runtime"
)

// ListImages lists existing images.
func (s *Server) ListImages(ctx context.Context, req *pb.ListImagesRequest) (*pb.ListImagesResponse, error) {
	logrus.Debugf("ListImages: %+v", req)
	filter := ""
	reqFilter := req.GetFilter()
	if reqFilter != nil {
		filterImage := reqFilter.GetImage()
		if filterImage != nil {
			filter = filterImage.GetImage()
		}
	}
	results, err := s.images.ListImages(filter)
	if err != nil {
		return nil, err
	}
	response := pb.ListImagesResponse{}
	for _, result := range results {
		response.Images = append(response.Images, &pb.Image{
			Id:       sPtr(result.ID),
			RepoTags: result.Names,
			Size_:    result.Size,
		})
	}
	return &response, nil
}
