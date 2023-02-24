package service

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizangela"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizbinah"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

type LibrarianSephirahServiceService struct {
	pb.UnimplementedLibrarianSephirahServiceServer

	t *biztiphereth.Tiphereth
	g *bizgebura.Gebura
	b *bizbinah.Binah
	y *bizyesod.Yesod
}

func NewLibrarianSephirahServiceService(
	a *bizangela.Angela,
	t *biztiphereth.Tiphereth,
	g *bizgebura.Gebura,
	b *bizbinah.Binah,
	y *bizyesod.Yesod,
) pb.LibrarianSephirahServiceServer {
	return &LibrarianSephirahServiceService{
		UnimplementedLibrarianSephirahServiceServer: pb.UnimplementedLibrarianSephirahServiceServer{},
		t: t,
		g: g,
		b: b,
		y: y,
	}
}
