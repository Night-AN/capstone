import { gql } from 'apollo-angular';

export const GET_ROLES = gql`
  query GetRoles($after: Cursor, $before: Cursor, $first: Int, $last: Int, $where: RoleWhereInput) {
    roles(after: $after, before: $before, first: $first, last: $last, where: $where) {
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

export const GET_ROLE_BY_ID = gql`
  query GetRoleById($id: UUID!) {
    node(id: $id) {
      ... on Role {
        id
        name
        description
        createdAt
        updatedAt
      }
    }
  }
`;

export const CREATE_ROLE = gql`
  mutation CreateRole($input: CreateRoleInput!) {
    createRole(input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const UPDATE_ROLE = gql`
  mutation UpdateRole($id: UUID!, $input: UpdateRoleInput!) {
    updateRole(id: $id, input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const DELETE_ROLE = gql`
  mutation DeleteRole($id: UUID!) {
    deleteRole(id: $id) {
      id
    }
  }
`;
