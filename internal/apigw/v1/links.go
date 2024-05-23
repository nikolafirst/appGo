package v1

import (
	"appGo/pkg/api/apiv1"
	"appGo/pkg/conv"
	"appGo/pkg/httputil"
	"appGo/pkg/pb"
	"net/http"
)

func NewLinksHandler(linksClient linksClient) *LinksHandler {
	return &LinksHandler{client: linksClient}
}

type LinksHandler struct {
	client linksClient
}

func (h *LinksHandler) GetLinks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resp, err := h.client.ListLinks(ctx, nil)
	if err != nil {
		handleGRPCError(w, err)
		return
	}

	linkList := make([]apiv1.Link, 0, len(resp.Links))
	for _, l := range resp.Links {
		linkList = append(
			linkList, apiv1.Link{
				CreatedAt: l.CreatedAt,
				Id:        l.Id,
				Images:    l.Images,
				Tags:      l.Tags,
				Title:     l.Title,
				UpdatedAt: l.UpdatedAt,
				Url:       l.Url,
				UserId:    l.UserId,
			},
		)
	}

	httputil.MarshalResponse(w, http.StatusOK, linkList)
}

func (h *LinksHandler) PostLinks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var l apiv1.LinkCreate
	code, err := httputil.Unmarshal(w, r, &l)
	if err != nil {
		httputil.MarshalResponse(
			w, code, apiv1.Error{
				Code:    httputil.ConvertHTTPToErrorCode(code),
				Message: conv.ToPtr(err.Error()),
			},
		)
		return
	}

	if _, err := h.client.CreateLink(
		ctx, &pb.CreateLinkRequest{
			Id:     l.Id,
			Title:  l.Title,
			Url:    l.Url,
			Images: l.Images,
			Tags:   l.Tags,
			UserId: l.UserId,
		},
	); err != nil {
		handleGRPCError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *LinksHandler) DeleteLinksId(w http.ResponseWriter, r *http.Request, id string) {
	ctx := r.Context()

	if _, err := h.client.DeleteLink(ctx, &pb.DeleteLinkRequest{Id: id}); err != nil {
		handleGRPCError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *LinksHandler) GetLinksId(w http.ResponseWriter, r *http.Request, id string) {
	ctx := r.Context()

	link, err := h.client.GetLink(ctx, &pb.GetLinkRequest{Id: id})
	if err != nil {
		handleGRPCError(w, err)
		return
	}

	httputil.MarshalResponse(
		w, http.StatusOK, apiv1.Link{
			CreatedAt: link.CreatedAt,
			Id:        link.Id,
			Images:    link.Images,
			Tags:      link.Tags,
			Title:     link.Title,
			UpdatedAt: link.UpdatedAt,
			Url:       link.Url,
			UserId:    link.UserId,
		},
	)
}

func (h *LinksHandler) PutLinksId(w http.ResponseWriter, r *http.Request, id string) {
	ctx := r.Context()
	var l apiv1.LinkCreate
	code, err := httputil.Unmarshal(w, r, &l)
	if err != nil {
		httputil.MarshalResponse(
			w, code, apiv1.Error{
				Code:    httputil.ConvertHTTPToErrorCode(code),
				Message: conv.ToPtr(err.Error()),
			},
		)
		return
	}

	if _, err := h.client.UpdateLink(
		ctx, &pb.UpdateLinkRequest{
			Id:     l.Id,
			Title:  l.Title,
			Url:    l.Url,
			Images: l.Images,
			Tags:   l.Tags,
			UserId: l.UserId,
		},
	); err != nil {
		handleGRPCError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *LinksHandler) GetLinksUserUserID(w http.ResponseWriter, r *http.Request, userID string) {
	ctx := r.Context()
	resp, err := h.client.GetLinkByUserID(ctx, &pb.GetLinksByUserId{UserId: userID})
	if err != nil {
		handleGRPCError(w, err)
		return
	}

	linkList := make([]apiv1.Link, 0, len(resp.Links))
	for _, l := range resp.Links {
		linkList = append(
			linkList, apiv1.Link{
				CreatedAt: l.CreatedAt,
				Id:        l.Id,
				Images:    l.Images,
				Tags:      l.Tags,
				Title:     l.Title,
				UpdatedAt: l.UpdatedAt,
				Url:       l.Url,
				UserId:    l.UserId,
			},
		)
	}

	httputil.MarshalResponse(w, http.StatusOK, linkList)
}
