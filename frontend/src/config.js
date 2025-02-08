export default {
  ipifyUrl: "https://api.ipify.org?format=json",
  backendUrl: (import.meta.env && import.meta.env.VITE_BACKEND_URL) || "http://localhost:8000/api/userinfo",
};