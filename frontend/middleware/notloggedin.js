export default function ({ $auth, redirect }) {
  const loggedIn = $auth.loggedIn;
  const user = $auth.user;
  if (loggedIn || user) return redirect("/");
}
