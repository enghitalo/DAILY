import { User } from "./user.entities";

export class UserService {
  constructor() {}
  getUserByIdService = getUserByIdService;
  addUserService = addUserService;
  getAllUserService = getAllUserService;
  getUserByUsernameService = getUserByUsernameService;
  getRemoveUserService = getRemoveUserService;
}

function getUserByIdService(id: string): User {
  var user: User = {
    id: 0,
    username: "",
    password: "",
    name: "",
    created_at: "",
    updated_at: "",
    deleted_at: "",
    active: false,
  };
  return user;
}

function addUserService(username: string, password: string): Boolean {
  var user: User = {
    id: 0,
    username: "",
    password: "",
    name: "",
    created_at: "",
    updated_at: "",
    deleted_at: "",
    active: false,
  };
  return true;
}

function getAllUserService(id: string): User[] {
  var user: User = {
    id: 0,
    username: "",
    password: "",
    name: "",
    created_at: "",
    updated_at: "",
    deleted_at: "",
    active: false,
  };
  return [user];
}

function getUserByUsernameService(id: string): User {
  var user: User = {
    id: 0,
    username: "",
    password: "",
    name: "",
    created_at: "",
    updated_at: "",
    deleted_at: "",
    active: false,
  };
  return user;
}

function getRemoveUserService(id: string): Boolean {
  var user: User = {
    id: 0,
    username: "",
    password: "",
    name: "",
    created_at: "",
    updated_at: "",
    deleted_at: "",
    active: false,
  };
  return true;
}
