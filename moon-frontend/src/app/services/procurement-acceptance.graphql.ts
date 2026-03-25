import { gql } from 'apollo-angular';

export const GET_PROCUREMENT_ACCEPTANCES = gql`
  query GetProcurementAcceptances($after: Cursor, $before: Cursor, $first: Int, $last: Int, $where: ProcurementAcceptanceWhereInput) {
    procurementAcceptances(after: $after, before: $before, first: $first, last: $last, where: $where) {
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

export const GET_PROCUREMENT_ACCEPTANCE_BY_ID = gql`
  query GetProcurementAcceptanceById($id: UUID!) {
    node(id: $id) {
      ... on ProcurementAcceptance {
        id
        name
        description
        createdAt
        updatedAt
      }
    }
  }
`;

export const CREATE_PROCUREMENT_ACCEPTANCE = gql`
  mutation CreateProcurementAcceptance($input: CreateProcurementAcceptanceInput!) {
    createProcurementAcceptance(input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const UPDATE_PROCUREMENT_ACCEPTANCE = gql`
  mutation UpdateProcurementAcceptance($id: UUID!, $input: UpdateProcurementAcceptanceInput!) {
    updateProcurementAcceptance(id: $id, input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const DELETE_PROCUREMENT_ACCEPTANCE = gql`
  mutation DeleteProcurementAcceptance($id: UUID!) {
    deleteProcurementAcceptance(id: $id) {
      id
    }
  }
`;
