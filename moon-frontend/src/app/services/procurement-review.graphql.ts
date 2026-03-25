import { gql } from 'apollo-angular';

export const GET_PROCUREMENT_REVIEWS = gql`
  query GetProcurementReviews($after: Cursor, $before: Cursor, $first: Int, $last: Int, $where: ProcurementReviewWhereInput) {
    procurementReviews(after: $after, before: $before, first: $first, last: $last, where: $where) {
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

export const GET_PROCUREMENT_REVIEW_BY_ID = gql`
  query GetProcurementReviewById($id: UUID!) {
    node(id: $id) {
      ... on ProcurementReview {
        id
        name
        description
        createdAt
        updatedAt
      }
    }
  }
`;

export const CREATE_PROCUREMENT_REVIEW = gql`
  mutation CreateProcurementReview($input: CreateProcurementReviewInput!) {
    createProcurementReview(input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const UPDATE_PROCUREMENT_REVIEW = gql`
  mutation UpdateProcurementReview($id: UUID!, $input: UpdateProcurementReviewInput!) {
    updateProcurementReview(id: $id, input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const DELETE_PROCUREMENT_REVIEW = gql`
  mutation DeleteProcurementReview($id: UUID!) {
    deleteProcurementReview(id: $id) {
      id
    }
  }
`;
