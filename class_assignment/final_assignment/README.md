# Social Media Post Handler

## Overview

This program demonstrates the use of concurrency patterns in Go to handle social media posts across multiple platforms. It leverages the fanout pattern to post content simultaneously to various social media platforms and the fanin pattern to aggregate engagement metrics such as comments and likes.

## Features

- **Create and Share Post**: Posts content to multiple social media platforms concurrently.
- **Aggregate Engagement**: Collects comments and likes from all platforms concurrently.

## Code Explanation

### Structures and Interfaces

#### Post

```go
type Post struct {
    title     string
    content   string
    mediaLink string
}
```

#### SocialMediaAPI

```go
type SocialMediaAPI interface {
    CreatePost(post Post) string
    FetchCommentsAndLikes(postID string) (comments int, likes int)
}
```

#### API

```go
type API struct {
    platform string
}
```

### Function

#### *createDummyAPI*
- Creates a list of dummy social media API implementations for different platforms.

#### *SharePostToAll*
- Orchestrates the sharing of a post to all platforms and waits for all operations to complete.
- It uses FanOut Patern

#### *aggregateEngagement*
- Aggregates comments and likes from all platforms for a given post ID.
- It uses FanIn Patern

### Exampl Output

```
all account-> [{{instagram}} {{twitter}} {{linkedin}} {{facebook}}]
instagram post title created
twitter post title created
linkedin post title created
facebook post title created
Post share to all account is Completed using FanOut pattern
waiting for peoples to like and comment XD


post id 1964
   total comments: 49
   total likes: 130
Collected all likes and comments using FanIn pattern
```