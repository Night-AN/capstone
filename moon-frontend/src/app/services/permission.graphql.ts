import { gql } from 'apollo-angular';

export const GET_PERMISSIONS = gql`
  query GetPermissions($after: Cursor, $before: Cursor, $first: Int, $last: Int, $where: PermissionWhereInput) {
    permissions(after: $after, before: $before, first: $first, last: $last, where: $where) {
      edges {
        cursor
        node {
          id
          name
          description
          createdAt
          updatedAt
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

export const GET_PERMISSION_BY_ID = gql`
  query GetPermissionById($id: UUID!) {
    node(id: $id) {
      ... on Permission {
        id
        name
        description
        createdAt
        updatedAt
      }
    }
  }
`;

export const CREATE_PERMISSION = gql`
  mutation CreatePermission($input: CreatePermissionInput!) {
    createPermission(input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const UPDATE_PERMISSION = gql`
  mutation UpdatePermission($id: UUID!, $input: UpdatePermissionInput!) {
    updatePermission(id: $id, input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const DELETE_PERMISSION = gql`
  mutation DeletePermission($id: UUID!) {
    deletePermission(id: $id) {
      id
    }
  }
`;
