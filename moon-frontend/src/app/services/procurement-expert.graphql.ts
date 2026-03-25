import { gql } from 'apollo-angular';

export const GET_PROCUREMENT_EXPERTS = gql`
  query GetProcurementExperts($after: Cursor, $before: Cursor, $first: Int, $last: Int, $where: ProcurementExpertWhereInput) {
    procurementExperts(after: $after, before: $before, first: $first, last: $last, where: $where) {
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

export const GET_PROCUREMENT_EXPERT_BY_ID = gql`
  query GetProcurementExpertById($id: UUID!) {
    node(id: $id) {
      ... on ProcurementExpert {
        id
        name
        description
        createdAt
        updatedAt
      }
    }
  }
`;

export const CREATE_PROCUREMENT_EXPERT = gql`
  mutation CreateProcurementExpert($input: CreateProcurementExpertInput!) {
    createProcurementExpert(input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const UPDATE_PROCUREMENT_EXPERT = gql`
  mutation UpdateProcurementExpert($id: UUID!, $input: UpdateProcurementExpertInput!) {
    updateProcurementExpert(id: $id, input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const DELETE_PROCUREMENT_EXPERT = gql`
  mutation DeleteProcurementExpert($id: UUID!) {
    deleteProcurementExpert(id: $id) {
      id
    }
  }
`;
