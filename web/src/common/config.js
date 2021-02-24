const WS_SCHEMA = location.protocol == "http:" ? "ws:" : "wss:";
export const API_URL = location.protocol + "//"+ location.hostname+":"+ location.port +"/";
export const WS_URL = WS_SCHEMA + "//" + location.hostname +":"+ location.port+"/v1/chat"
export default { API_URL, WS_URL};