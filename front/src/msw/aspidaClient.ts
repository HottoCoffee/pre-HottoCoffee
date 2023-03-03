import aspida from "@aspida/axios";
import api from "../../api/$api";
export const client = api(aspida());
