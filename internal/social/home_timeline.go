package social

import (
	"context"
	"github.com/DKW2/MuCache_Extended/pkg/invoke"
	"github.com/DKW2/MuCache_Extended/pkg/state"
)

func ReadHomeTimeline(ctx context.Context, userId string) []Post {
	postIds, err := state.GetState[[]string](ctx, userId)
	if err != nil {
		return []Post{}
	}
	req := ReadPostsRequest{PostIds: postIds}
	postsResp := invoke.Invoke[ReadPostsResponse](ctx, "poststorage", "ro_read_posts", req)
	return postsResp.Posts
}

func WriteHomeTimeline(ctx context.Context, creatorId string, newPostIds []string) {
	req := GetFollowersRequest{UserId: creatorId}
	resp := invoke.Invoke[GetFollowersResponse](ctx, "socialgraph", "ro_get_followers", req)
	for _, follower := range resp.Followers {
		postIds, err := state.GetState[[]string](ctx, follower)
		if err != nil {
			postIds = []string{}
		}
		if len(postIds) >= 10 {
			postIds = postIds[1:]
		}
		postIds = append(postIds, newPostIds...)
		state.SetState(ctx, follower, postIds)
	}
}
