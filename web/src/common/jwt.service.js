const ID_TOKEN_KEY = "token";
const ID_ROLE = "role";

export const getToken = () => {
  return window.localStorage.getItem(ID_TOKEN_KEY);
};

export const getRole = () => {
  return window.localStorage.getItem(ID_ROLE);
};

export const saveToken = token => {
  window.localStorage.setItem(ID_TOKEN_KEY, token);
};

export const setRole = role => {
  window.localStorage.setItem(ID_ROLE, role);
};

export const destroyToken = () => {
  window.localStorage.removeItem(ID_TOKEN_KEY);
  window.localStorage.removeItem(ID_ROLE);
};

export default { getToken, saveToken, destroyToken, setRole, getRole };