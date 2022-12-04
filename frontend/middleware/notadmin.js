export default function ({ $auth, redirect }) {
  const user = $auth?.user;
  if (user?.role === "admin") return redirect("/");
}
