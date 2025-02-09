export default {
  ipifyUrl: "https://api.ipify.org?format=json",
  backendUrl: process.env.VUE_APP_BACKEND_URL || "http://localhost:8000/api/userinfo",
};
