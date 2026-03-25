import { gql } from 'apollo-angular';

// 查询文件列表
export const GET_FILES = gql`
  query Files($first: Int, $after: Cursor, $where: FileWhereInput) {
    files(first: $first, after: $after, where: $where) {
      edges {
        node {
          id
          fileName
          fileType
          fileSize
          createdAt
          updatedAt
        }
        cursor
      }
      pageInfo {
        hasNextPage
        hasPreviousPage
        startCursor
        endCursor
      }
      totalCount
    }
  }
`;

// 根据ID查询文件
export const GET_FILE_BY_ID = gql`
  query File($id: ID!) {
    file(id: $id) {
      id
      fileName
      fileType
      fileSize
      createdAt
      updatedAt
    }
  }
`;

// 创建文件
export const CREATE_FILE = gql`
  mutation CreateFile($input: CreateFileInput!) {
    createFile(input: $input) {
      id
      fileName
      fileType
      fileSize
      createdAt
      updatedAt
    }
  }
`;

// 更新文件
export const UPDATE_FILE = gql`
  mutation UpdateFile($id: ID!, $input: UpdateFileInput!) {
    updateFile(id: $id, input: $input) {
      id
      fileName
      fileType
      fileSize
      createdAt
      updatedAt
    }
  }
`;

// 删除文件
export const DELETE_FILE = gql`
  mutation DeleteFile($id: ID!) {
    deleteFile(id: $id) {
      id
      fileName
    }
  }
`;