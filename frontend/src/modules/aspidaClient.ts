import aspida from "@aspida/axios";
import api from "../api/$api";

/** @package */
export const client = api(aspida());
