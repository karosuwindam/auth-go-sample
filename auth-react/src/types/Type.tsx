export const RoleType = {
    Admin: 'admin',
    User: 'user',
    Guest: 'guest'
}

export type Roles = typeof RoleType[keyof typeof RoleType];

export const SetRole = (roleData:Roles) => {
    localStorage.setItem('user',roleData)
}