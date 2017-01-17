package server

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	pb "k8s.io/kubernetes/pkg/kubelet/api/v1alpha1/runtime"
)

// RemoveImage removes the image.
func (s *Server) RemoveImage(ctx context.Context, req *pb.RemoveImageRequest) (*pb.RemoveImageResponse, error) {
	logrus.Debugf("RemoveImage: %+v", req)
	image := ""
	img := req.GetImage()
	if img != nil {
		image = img.GetImage()
	}
	if image == "" {
		return nil, fmt.Errorf("no image specified")
	}
	err := s.images.RemoveImage(s.imageContext, image)
	if err != nil {
		return nil, err
	}
	return &pb.RemoveImageResponse{}, nil
}
