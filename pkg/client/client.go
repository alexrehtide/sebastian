package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/alexrehtide/sebastian/api/v1"
)

func NewClient() *Client {
	return &Client{}
}

type Client struct {
	pbClient pb.SebastianClient
	conn     *grpc.ClientConn
}

func (c *Client) Connect(addr string) error {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("did not connect: %w", err)
	}
	c.conn = conn

	c.pbClient = pb.NewSebastianClient(conn)
	return nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Authenticate(ctx context.Context, req AuthenticateRequest) (AuthenticateResponse, error) {
	r, err := c.pbClient.Authenticate(ctx, &pb.AuthenticateRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		return AuthenticateResponse{}, fmt.Errorf("Client.Authenticate: %w", err)
	}
	return AuthenticateResponse{
		AccessToken:  r.GetAccessToken(),
		RefreshToken: r.GetRefreshToken(),
	}, nil
}

type AuthenticateRequest struct {
	Email    string
	Password string
}

type AuthenticateResponse struct {
	AccessToken  string
	RefreshToken string
}
