import { gql } from 'apollo-angular';

export const GET_ASSET_TYPES = gql`
  query GetAssetTypes($after: Cursor, $before: Cursor, $first: Int, $last: Int, $where: AssetTypeWhereInput) {
    assetTypes(after: $after, before: $before, first: $first, last: $last, where: $where) {
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

export const GET_ASSET_TYPE_BY_ID = gql`
  query GetAssetTypeById($id: UUID!) {
    node(id: $id) {
      ... on AssetType {
        id
        name
        description
        createdAt
        updatedAt
      }
    }
  }
`;

export const CREATE_ASSET_TYPE = gql`
  mutation CreateAssetType($input: CreateAssetTypeInput!) {
    createAssetType(input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const UPDATE_ASSET_TYPE = gql`
  mutation UpdateAssetType($id: UUID!, $input: UpdateAssetTypeInput!) {
    updateAssetType(id: $id, input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const DELETE_ASSET_TYPE = gql`
  mutation DeleteAssetType($id: UUID!) {
    deleteAssetType(id: $id) {
      id
    }
  }
`;
