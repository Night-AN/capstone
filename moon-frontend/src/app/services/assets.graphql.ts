import { gql } from 'apollo-angular';

export const GET_ASSETS = gql`
  query GetAssets($after: Cursor, $before: Cursor, $first: Int, $last: Int, $where: AssetsWhereInput) {
    assets(after: $after, before: $before, first: $first, last: $last, where: $where) {
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

export const GET_ASSET_BY_ID = gql`
  query GetAssetById($id: UUID!) {
    node(id: $id) {
      ... on Assets {
        id
        name
        description
        createdAt
        updatedAt
      }
    }
  }
`;

export const CREATE_ASSET = gql`
  mutation CreateAsset($input: CreateAssetsInput!) {
    createAsset(input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const UPDATE_ASSET = gql`
  mutation UpdateAsset($id: UUID!, $input: UpdateAssetsInput!) {
    updateAsset(id: $id, input: $input) {
      id
      name
      description
      createdAt
      updatedAt
    }
  }
`;

export const DELETE_ASSET = gql`
  mutation DeleteAsset($id: UUID!) {
    deleteAsset(id: $id) {
      id
    }
  }
`;
