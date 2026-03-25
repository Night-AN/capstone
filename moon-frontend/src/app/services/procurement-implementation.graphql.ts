import { gql } from 'apollo-angular';

export const GET_PROCUREMENT_IMPLEMENTATIONS = gql`
  query GetProcurementImplementations($after: Cursor, $before: Cursor, $first: Int, $last: Int, $where: ProcurementImplementationWhereInput) {
    procurementImplementations(after: $after, before: $before, first: $first, last: $last, where: $where) {
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

export const GET_PROCUREMENT_IMPLEMENTATION_BY_ID = gql`
  query GetProcurementImplementationById($id: UUID!) {
    node(id: $id) {
      ... on ProcurementImplementation {
        id
        name
        description
        createdAt
        updatedAt
      }
    }
  }
`;

export const CREATE_PROCUREMENT_IMPLEMENTATION = gql`
  mutation CreateProcurementImplementation($input: CreateProcurementImplementationInput!) {
    createProcurementImplementation(input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const UPDATE_PROCUREMENT_IMPLEMENTATION = gql`
  mutation UpdateProcurementImplementation($id: UUID!, $input: UpdateProcurementImplementationInput!) {
    updateProcurementImplementation(id: $id, input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const DELETE_PROCUREMENT_IMPLEMENTATION = gql`
  mutation DeleteProcurementImplementation($id: UUID!) {
    deleteProcurementImplementation(id: $id) {
      id
    }
  }
`;
