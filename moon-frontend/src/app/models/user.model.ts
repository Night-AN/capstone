export interface User {
  user_id: string;
  nickname: string;
  full_name: string;
  email: string;
  created_at: string;
  updated_at: string;
}

export interface UserCreateRequest {
  nickname: string;
  full_name: string;
  email: string;
  password: string;
}

export interface UserUpdateRequest {
  user_id: string;
  nickname: string;
  full_name: string;
  email: string;
}

export interface UserDeleteRequest {
  user_id: string;
}

export interface UserListResponse {
  code: string;
  message: string;
  data: {
    users: UserListItem[];
  };
}

export interface UserGetResponse {
  code: string;
  message: string;
  data: {
    user_id: string;
    nickname: string;
    full_name: string;
    email: string;
  };
}

export interface UserListItem {
  user_id: string;
  nickname: string;
  full_name: string;
  email: string;
}
