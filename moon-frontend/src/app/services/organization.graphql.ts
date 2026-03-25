import { gql } from 'apollo-angular';

export const GET_ORGANIZATIONS = gql`
  query GetOrganizations($after: Cursor, $before: Cursor, $first: Int, $last: Int, $where: OrganizationWhereInput) {
    organizations(after: $after, before: $before, first: $first, last: $last, where: $where) {
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

export const GET_ORGANIZATION_BY_ID = gql`
  query GetOrganizationById($id: UUID!) {
    node(id: $id) {
      ... on Organization {
        id
        name
        description
        createdAt
        updatedAt
      }
    }
  }
`;

export const CREATE_ORGANIZATION = gql`
  mutation CreateOrganization($input: CreateOrganizationInput!) {
    createOrganization(input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const UPDATE_ORGANIZATION = gql`
  mutation UpdateOrganization($id: UUID!, $input: UpdateOrganizationInput!) {
    updateOrganization(id: $id, input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const DELETE_ORGANIZATION = gql`
  mutation DeleteOrganization($id: UUID!) {
    deleteOrganization(id: $id) {
      id
    }
  }
`;
