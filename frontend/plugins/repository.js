import createRepository from "~/repositories/index";

export default (ctx, inject) => {
  inject("repositories", createRepository(ctx.$axios));
};
