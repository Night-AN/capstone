import { gql } from 'apollo-angular';

export const GET_PROCUREMENT_PLAN_TYPES = gql`
  query GetProcurementPlanTypes($after: Cursor, $before: Cursor, $first: Int, $last: Int, $where: ProcurementPlanTypeWhereInput) {
    procurementPlanTypes(after: $after, before: $before, first: $first, last: $last, where: $where) {
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

export const GET_PROCUREMENT_PLAN_TYPE_BY_ID = gql`
  query GetProcurementPlanTypeById($id: UUID!) {
    node(id: $id) {
      ... on ProcurementPlanType {
        id
        name
        description
        createdAt
        updatedAt
      }
    }
  }
`;

export const CREATE_PROCUREMENT_PLAN_TYPE = gql`
  mutation CreateProcurementPlanType($input: CreateProcurementPlanTypeInput!) {
    createProcurementPlanType(input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const UPDATE_PROCUREMENT_PLAN_TYPE = gql`
  mutation UpdateProcurementPlanType($id: UUID!, $input: UpdateProcurementPlanTypeInput!) {
    updateProcurementPlanType(id: $id, input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const DELETE_PROCUREMENT_PLAN_TYPE = gql`
  mutation DeleteProcurementPlanType($id: UUID!) {
    deleteProcurementPlanType(id: $id) {
      id
    }
  }
`;
