import Users from "./UsersRepository";

export default ($axios) => ({
  users: Users($axios),
});
