export default function ({ $auth, redirect }) {
  const user = $auth?.user;
  if (user?.role !== "Admin") return redirect("/");
}
