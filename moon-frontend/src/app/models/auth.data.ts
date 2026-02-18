
export interface User {
    UserID:string
    Username:string
    Organization:string
    Role:string[]
}

export interface AuthResponse {
  token: string;
  user: User;
}

export interface LoginCredentials{
  email: string;
  password: string;
}

export interface RegisterData{
  nickname: string;
  fullName: string;
  email: string;
  password: string;
}