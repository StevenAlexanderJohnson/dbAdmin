import { writable } from "svelte/store";

export const selectedConnection = writable('');
export const selectedUser = writable('');