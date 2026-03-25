import { gql } from 'apollo-angular';

export const GET_USERS = gql`
  query GetUsers($after: Cursor, $before: Cursor, $first: Int, $last: Int, $where: UserWhereInput) {
    users(after: $after, before: $before, first: $first, last: $last, where: $where) {
      edges {
        cursor
        node {
          id
          email
          nickname
          fullname
          createdat
          updatedat
        }
      }
      pageInfo {
        endCursor
        hasNextPage
        hasPreviousPage
        startCursor
      }
    }
  }
`;

export const GET_USER_BY_ID = gql`
  query GetUserById($id: UUID!) {
    node(id: $id) {
      ... on User {
        id
        email
        name
        createdAt
        updatedAt
      }
    }
  }
`;

export const CREATE_USER = gql`
  mutation CreateUser($input: CreateUserInput!) {
    createUser(input: $input) {
      nickname
    	fullname
      email
			passwordhash
    }
  }
`;

export const UPDATE_USER = gql`
  mutation UpdateUser($id: UUID!, $input: UpdateUserInput!) {
    updateUser(id: $id, input: $input) {
      id
      email
      name
      createdAt
      updatedAt
    }
  }
`;

export const DELETE_USER = gql`
  mutation DeleteUser($id: UUID!) {
    deleteUser(id: $id) {
      id
    }
  }
`;
