import { gql } from 'apollo-angular';

export const GET_PROCUREMENT_PLANS = gql`
  query GetProcurementPlans($after: Cursor, $before: Cursor, $first: Int, $last: Int, $where: ProcurementPlanWhereInput) {
    procurementPlans(after: $after, before: $before, first: $first, last: $last, where: $where) {
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

export const GET_PROCUREMENT_PLAN_BY_ID = gql`
  query GetProcurementPlanById($id: UUID!) {
    node(id: $id) {
      ... on ProcurementPlan {
        id
        name
        description
        createdAt
        updatedAt
      }
    }
  }
`;

export const CREATE_PROCUREMENT_PLAN = gql`
  mutation CreateProcurementPlan($input: CreateProcurementPlanInput!) {
    createProcurementPlan(input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const UPDATE_PROCUREMENT_PLAN = gql`
  mutation UpdateProcurementPlan($id: UUID!, $input: UpdateProcurementPlanInput!) {
    updateProcurementPlan(id: $id, input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const DELETE_PROCUREMENT_PLAN = gql`
  mutation DeleteProcurementPlan($id: UUID!) {
    deleteProcurementPlan(id: $id) {
      id
    }
  }
`;
