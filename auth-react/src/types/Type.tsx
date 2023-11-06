export const RoleType = {
    Admin: 'admin',
    User: 'user',
    Guest: 'guest'
}

export type Roles = typeof RoleType[keyof typeof RoleType];

export const UserType = {
    name: '',
    role: ''
}

export type User = typeof UserType[keyof typeof UserType];

export const SetRole = (roleData:Roles) => {
    localStorage.setItem('user',roleData)
}

export const GetRole = () => {
    const user = {
        name: sessionStorage.getItem('user') || '',
        role: sessionStorage.getItem('role') || ''
    }
    return user;
}
