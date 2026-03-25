import { gql } from 'apollo-angular';

export const GET_ASSET_CATEGORIES = gql`
  query GetAssetCategories($after: Cursor, $before: Cursor, $first: Int, $last: Int, $where: AssetCategoryWhereInput) {
    assetCategories(after: $after, before: $before, first: $first, last: $last, where: $where) {
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

export const GET_ASSET_CATEGORY_BY_ID = gql`
  query GetAssetCategoryById($id: UUID!) {
    node(id: $id) {
      ... on AssetCategory {
        id
        name
        description
        createdAt
        updatedAt
      }
    }
  }
`;

export const CREATE_ASSET_CATEGORY = gql`
  mutation CreateAssetCategory($input: CreateAssetCategoryInput!) {
    createAssetCategory(input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const UPDATE_ASSET_CATEGORY = gql`
  mutation UpdateAssetCategory($id: UUID!, $input: UpdateAssetCategoryInput!) {
    updateAssetCategory(id: $id, input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const DELETE_ASSET_CATEGORY = gql`
  mutation DeleteAssetCategory($id: UUID!) {
    deleteAssetCategory(id: $id) {
      id
    }
  }
`;
