# API Reference

Base URL: configured via `API_BASE_URL` in frontend `.env`.

All protected endpoints require `Authorization: Bearer <access_token>` header.
All request/response bodies are JSON unless noted otherwise.

---

## Authentication

### POST /auth/register
Create a new user account.

**Request**
```json
{
  "email": "user@example.com",
  "password": "secret"
}
```

**Response** `201`
```json
{
  "access_token": "...",
  "refresh_token": "...",
  "expires_in": 900,
  "user": { "id": 1, "email": "user@example.com" }
}
```

---

### POST /auth/login
Authenticate and receive tokens.

**Request**
```json
{
  "email": "user@example.com",
  "password": "secret"
}
```

**Response** `200`
```json
{
  "access_token": "...",
  "refresh_token": "...",
  "expires_in": 900,
  "user": { "id": 1, "email": "user@example.com" }
}
```

---

### POST /auth/refresh
Exchange a refresh token for a new token pair.

**Request**
```json
{
  "refresh_token": "..."
}
```

**Response** `200`
```json
{
  "access_token": "...",
  "refresh_token": "...",
  "expires_in": 900,
  "user": { "id": 1, "email": "user@example.com" }
}
```

---

### GET /auth/session
🔒 Get the current user's info.

**Response** `200`
```json
{
  "user": { "id": 1, "email": "user@example.com" }
}
```

---

## Profiles

### GET /profiles/
🔒 Get all profiles belonging to the authenticated user.

**Response** `200` — array of profile objects (see shape below).

---

### POST /profiles/
🔒 Create a new profile.

**Request**
```json
{
  "nickname": "my-profile",
  "first_name": "Jane",
  "last_name": "Doe",
  "email": "jane@example.com",
  "avatar_url": "https://..."
}
```
- `nickname` required, 2–255 chars
- `email` and `avatar_url` optional but validated if present

**Response** `201` — created profile object.

---

### GET /profiles/{id}
🔒 Get a single profile by ID (must belong to the authenticated user).

**Response** `200` — profile object.

---

### PUT /profiles/{id}
🔒 Update a profile.

**Request**
```json
{
  "first_name": "Jane",
  "last_name": "Doe",
  "email": "jane@example.com",
  "avatar_url": "https://..."
}
```
All fields optional. `avatar_url` is preserved from the existing record if omitted.

**Response** `200` — updated profile object.

---

### POST /profiles/{id}/avatar
🔒 Upload an avatar image for a profile.

**Request** `multipart/form-data`
- Field: `avatar` — image file (PNG, JPEG, or BMP; max 2 MB)

**Response** `200`
```json
{
  "avatar_url": "/uploads/avatars/filename.png"
}
```

---

### DELETE /profiles/{id}
🔒 Delete a profile.

**Response** `204 No Content`

---

### GET /public/profiles/{id}
Public. Get a profile by ID (no auth required — used for shareable links).

**Response** `200` — profile object (same shape as protected endpoints).

---

#### Profile object shape
```json
{
  "id": 1,
  "user_id": 1,
  "nickname": "my-profile",
  "first_name": "Jane",
  "last_name": "Doe",
  "email": "jane@example.com",
  "avatar_url": "/uploads/avatars/abc.png",
  "created_at": "2025-01-01T00:00:00Z",
  "updated_at": null,
  "links": [
    {
      "id": 1,
      "profile_id": 1,
      "platform_id": 2,
      "url": "https://github.com/jane",
      "position": 0,
      "platform": {
        "id": 2,
        "name": "GitHub",
        "icon": "github",
        "color": "#181717"
      },
      "created_at": "2025-01-01T00:00:00Z",
      "updated_at": null
    }
  ]
}
```

---

## Links

### GET /links/{profile_id}
🔒 Get all links for a profile.

**Response** `200` — array of link objects (see profile shape above for link structure).

---

### PUT /links/{profile_id}
🔒 Replace all links for a profile. Send the full desired list; existing links not in the list are deleted.

**Request**
```json
[
  {
    "platform_id": 2,
    "url": "https://github.com/jane",
    "position": 0
  }
]
```
- `id` may be omitted for new links
- `profile_id` in each item is overridden by the URL param

**Response** `200`

---

## Platforms

### GET /platforms/
🔒 Get all available platforms (GitHub, Twitter, etc.).

**Response** `200`
```json
[
  {
    "id": 1,
    "name": "GitHub",
    "icon": "github",
    "color": "#181717"
  }
]
```

---

## Static Files

### GET /uploads/*
Public. Serves uploaded files (e.g. avatars). Use the `avatar_url` value returned by the API directly as the path.
