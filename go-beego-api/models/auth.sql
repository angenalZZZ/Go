CREATE TABLE [dbo].[AuthOrg] (
	[Id] varchar(32) NOT NULL PRIMARY KEY CLUSTERED,
	[Code] varchar(32) NOT NULL,
	[Name] nvarchar(128) NOT NULL,
	[FullName] nvarchar(128) NOT NULL,
	[ShortName] nvarchar(128) NOT NULL,
	[SortCode] nvarchar(32) NOT NULL,
	[ParentId] varchar(32) NOT NULL,
	[Level] varchar(32) NOT NULL,
	[OrgType] varchar(32) NOT NULL,
	[Leader] varchar(32) NOT NULL,
	[Remark] nvarchar(512) NOT NULL,
	[Revision] int NOT NULL,
	[CreatedBy] varchar(32) NOT NULL,
	[CreatedTime] datetime NOT NULL,
	[UpdatedBy] varchar(32) NOT NULL,
	[UpdatedTime] datetime NOT NULL
) ON [PRIMARY]
GO
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'机构ID',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrg',@level2type=N'Column',@level2name=N'Id';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'机构代号',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrg',@level2type=N'Column',@level2name=N'Code';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'机构名',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrg',@level2type=N'Column',@level2name=N'Name';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'机构路径全称',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrg',@level2type=N'Column',@level2name=N'FullName';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'机构简称',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrg',@level2type=N'Column',@level2name=N'ShortName';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'排序代码',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrg',@level2type=N'Column',@level2name=N'SortCode';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'上级机构',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrg',@level2type=N'Column',@level2name=N'ParentId';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'机构级别',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrg',@level2type=N'Column',@level2name=N'Level';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'机构类型',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrg',@level2type=N'Column',@level2name=N'OrgType';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'负责人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrg',@level2type=N'Column',@level2name=N'Leader';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'机构说明',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrg',@level2type=N'Column',@level2name=N'Remark';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'乐观锁',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrg',@level2type=N'Column',@level2name=N'Revision';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrg',@level2type=N'Column',@level2name=N'CreatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrg',@level2type=N'Column',@level2name=N'CreatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrg',@level2type=N'Column',@level2name=N'UpdatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrg',@level2type=N'Column',@level2name=N'UpdatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'部门机构',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrg';
GO

CREATE TABLE [dbo].[AuthOrgProperty] (
	[Id] varchar(32) NOT NULL PRIMARY KEY CLUSTERED,
	[OrgId] varchar(32) NOT NULL,
	[Name] nvarchar(48) NOT NULL,
	[Value] nvarchar(1024) NOT NULL,
	[Revision] int NOT NULL,
	[CreatedBy] varchar(32) NOT NULL,
	[CreatedTime] datetime NOT NULL,
	[UpdatedBy] varchar(32) NOT NULL,
	[UpdatedTime] datetime NOT NULL
) ON [PRIMARY]
GO
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'属性ID',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrgProperty',@level2type=N'Column',@level2name=N'Id';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'机构ID',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrgProperty',@level2type=N'Column',@level2name=N'OrgId';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'属性名',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrgProperty',@level2type=N'Column',@level2name=N'Name';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'属性值',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrgProperty',@level2type=N'Column',@level2name=N'Value';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'乐观锁',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrgProperty',@level2type=N'Column',@level2name=N'Revision';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrgProperty',@level2type=N'Column',@level2name=N'CreatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrgProperty',@level2type=N'Column',@level2name=N'CreatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrgProperty',@level2type=N'Column',@level2name=N'UpdatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrgProperty',@level2type=N'Column',@level2name=N'UpdatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'机构属性',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthOrgProperty';
GO

CREATE TABLE [dbo].[AuthPermit] (
	[Code] varchar(128) NOT NULL PRIMARY KEY CLUSTERED,
	[Name] varchar(128) NOT NULL,
	[Revision] int NOT NULL,
	[CreatedBy] varchar(32) NOT NULL,
	[CreatedTime] datetime NOT NULL,
	[UpdatedBy] varchar(32) NOT NULL,
	[UpdatedTime] datetime NOT NULL
) ON [PRIMARY]
GO
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'权限代码',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthPermit',@level2type=N'Column',@level2name=N'Code';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'权限名称',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthPermit',@level2type=N'Column',@level2name=N'Name';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'乐观锁',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthPermit',@level2type=N'Column',@level2name=N'Revision';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthPermit',@level2type=N'Column',@level2name=N'CreatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthPermit',@level2type=N'Column',@level2name=N'CreatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthPermit',@level2type=N'Column',@level2name=N'UpdatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthPermit',@level2type=N'Column',@level2name=N'UpdatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'权限信息 自定义权限',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthPermit';
GO

CREATE TABLE [dbo].[AuthRole] (
	[Id] varchar(32) NOT NULL PRIMARY KEY CLUSTERED,
	[Code] varchar(32) NOT NULL,
	[SortCode] varchar(8) NOT NULL,
	[Name] nvarchar(48) NOT NULL,
	[Type] varchar(32) NOT NULL,
	[InWorkFlow] varchar(1) NOT NULL,
	[Status] varchar(32) NOT NULL,
	[Summary] nvarchar(512) NOT NULL,
	[Revision] int NOT NULL,
	[CreatedBy] varchar(32) NOT NULL,
	[CreatedTime] datetime NOT NULL,
	[UpdatedBy] varchar(32) NOT NULL,
	[UpdatedTime] datetime NOT NULL
) ON [PRIMARY]
GO
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'角色ID',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRole',@level2type=N'Column',@level2name=N'Id';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'角色代码',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRole',@level2type=N'Column',@level2name=N'Code';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'排序代码',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRole',@level2type=N'Column',@level2name=N'SortCode';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'角色名',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRole',@level2type=N'Column',@level2name=N'Name';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'角色类型',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRole',@level2type=N'Column',@level2name=N'Type';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'是否应用于工作流',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRole',@level2type=N'Column',@level2name=N'InWorkFlow';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'角色状态',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRole',@level2type=N'Column',@level2name=N'Status';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'角色描述',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRole',@level2type=N'Column',@level2name=N'Summary';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'乐观锁',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRole',@level2type=N'Column',@level2name=N'Revision';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRole',@level2type=N'Column',@level2name=N'CreatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRole',@level2type=N'Column',@level2name=N'CreatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRole',@level2type=N'Column',@level2name=N'UpdatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRole',@level2type=N'Column',@level2name=N'UpdatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'角色信息',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRole';
GO

CREATE TABLE [dbo].[AuthRolePermit] (
	[Id] varchar(32) NOT NULL PRIMARY KEY CLUSTERED,
	[PermitCode] varchar(128) NOT NULL,
	[RoleId] varchar(32) NOT NULL,
	[Revision] int NOT NULL,
	[CreatedBy] varchar(32) NOT NULL,
	[CreatedTime] datetime NOT NULL,
	[UpdatedBy] varchar(32) NOT NULL,
	[UpdatedTime] datetime NOT NULL
) ON [PRIMARY]
GO
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'角色权限ID',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRolePermit',@level2type=N'Column',@level2name=N'Id';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'权限代码',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRolePermit',@level2type=N'Column',@level2name=N'PermitCode';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'角色ID',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRolePermit',@level2type=N'Column',@level2name=N'RoleId';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'乐观锁',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRolePermit',@level2type=N'Column',@level2name=N'Revision';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRolePermit',@level2type=N'Column',@level2name=N'CreatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRolePermit',@level2type=N'Column',@level2name=N'CreatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRolePermit',@level2type=N'Column',@level2name=N'UpdatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRolePermit',@level2type=N'Column',@level2name=N'UpdatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'角色权限',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthRolePermit';
GO

CREATE TABLE [dbo].[AuthUser] (
	[Id] varchar(32) NOT NULL PRIMARY KEY CLUSTERED,
	[Code] varchar(32) NOT NULL,
	[Name] nvarchar(48) NOT NULL,
	[Password] varchar(32) NOT NULL,
	[Salt] varchar(24) NOT NULL,
	[Avatar] varchar(64) NOT NULL,
	[OrgId] varchar(32) NOT NULL,
	[Email] nvarchar(32) NOT NULL,
	[Phone] varchar(48) NOT NULL,
	[Status] varchar(32) NOT NULL,
	[Revision] int NOT NULL,
	[CreatedBy] varchar(32) NOT NULL,
	[CreatedTime] datetime NOT NULL,
	[UpdatedBy] varchar(32) NOT NULL,
	[UpdatedTime] datetime NOT NULL
) ON [PRIMARY]
GO
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'用户ID',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUser',@level2type=N'Column',@level2name=N'Id';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'用户代码',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUser',@level2type=N'Column',@level2name=N'Code';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'用户名',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUser',@level2type=N'Column',@level2name=N'Name';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'密码',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUser',@level2type=N'Column',@level2name=N'Password';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'密码盐值',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUser',@level2type=N'Column',@level2name=N'Salt';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'头像',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUser',@level2type=N'Column',@level2name=N'Avatar';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'机构',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUser',@level2type=N'Column',@level2name=N'OrgId';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'邮件',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUser',@level2type=N'Column',@level2name=N'Email';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'手机号',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUser',@level2type=N'Column',@level2name=N'Phone';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'状态',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUser',@level2type=N'Column',@level2name=N'Status';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'乐观锁',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUser',@level2type=N'Column',@level2name=N'Revision';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUser',@level2type=N'Column',@level2name=N'CreatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUser',@level2type=N'Column',@level2name=N'CreatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUser',@level2type=N'Column',@level2name=N'UpdatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUser',@level2type=N'Column',@level2name=N'UpdatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'用户信息',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUser';
GO

CREATE TABLE [dbo].[AuthUserAccount] (
	[Id] varchar(32) NOT NULL PRIMARY KEY CLUSTERED,
	[UserId] varchar(32) NOT NULL,
	[AccountCode] nvarchar(32) NOT NULL,
	[AccountType] nvarchar(32) NOT NULL,
	[Password] varchar(32) NOT NULL,
	[Status] varchar(32) NOT NULL,
	[Revision] int NOT NULL,
	[CreatedBy] varchar(32) NOT NULL,
	[CreatedTime] datetime NOT NULL,
	[UpdatedBy] varchar(32) NOT NULL,
	[UpdatedTime] datetime NOT NULL
) ON [PRIMARY]
GO
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'账号ID',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserAccount',@level2type=N'Column',@level2name=N'Id';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'用户ID',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserAccount',@level2type=N'Column',@level2name=N'UserId';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'账号代号',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserAccount',@level2type=N'Column',@level2name=N'AccountCode';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'账号类型',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserAccount',@level2type=N'Column',@level2name=N'AccountType';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'密码',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserAccount',@level2type=N'Column',@level2name=N'Password';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'账号状态',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserAccount',@level2type=N'Column',@level2name=N'Status';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'乐观锁',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserAccount',@level2type=N'Column',@level2name=N'Revision';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserAccount',@level2type=N'Column',@level2name=N'CreatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserAccount',@level2type=N'Column',@level2name=N'CreatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserAccount',@level2type=N'Column',@level2name=N'UpdatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserAccount',@level2type=N'Column',@level2name=N'UpdatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'用户信息',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserAccount';
GO

CREATE TABLE [dbo].[AuthUserBehavior] (
	[Id] varchar(32) NOT NULL PRIMARY KEY CLUSTERED,
	[UserId] varchar(32) NOT NULL,
	[ObjectId] nvarchar(32) NOT NULL,
	[ObjectType] nvarchar(32) NOT NULL,
	[Type] varchar(32) NOT NULL,
	[Value] varchar(32) NOT NULL,
	[Memo] varchar(512) NOT NULL,
	[Revision] int NOT NULL,
	[CreatedBy] varchar(32) NOT NULL,
	[CreatedTime] datetime NOT NULL,
	[UpdatedBy] varchar(32) NOT NULL,
	[UpdatedTime] datetime NOT NULL
) ON [PRIMARY]
GO
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'用户行为ID',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserBehavior',@level2type=N'Column',@level2name=N'Id';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'用户ID',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserBehavior',@level2type=N'Column',@level2name=N'UserId';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'关联对象ID',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserBehavior',@level2type=N'Column',@level2name=N'ObjectId';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'关联对象类型',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserBehavior',@level2type=N'Column',@level2name=N'ObjectType';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'行为类型',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserBehavior',@level2type=N'Column',@level2name=N'Type';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'行为数值',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserBehavior',@level2type=N'Column',@level2name=N'Value';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'行为说明',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserBehavior',@level2type=N'Column',@level2name=N'Memo';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'乐观锁',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserBehavior',@level2type=N'Column',@level2name=N'Revision';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserBehavior',@level2type=N'Column',@level2name=N'CreatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserBehavior',@level2type=N'Column',@level2name=N'CreatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserBehavior',@level2type=N'Column',@level2name=N'UpdatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserBehavior',@level2type=N'Column',@level2name=N'UpdatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'用户行为',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserBehavior';
GO

CREATE TABLE [dbo].[AuthUserPermit] (
	[Id] varchar(32) NOT NULL PRIMARY KEY CLUSTERED,
	[PermitCode] varchar(128) NOT NULL,
	[UserId] varchar(32) NOT NULL,
	[Revision] int NOT NULL,
	[CreatedBy] varchar(32) NOT NULL,
	[CreatedTime] datetime NOT NULL,
	[UpdatedBy] varchar(32) NOT NULL,
	[UpdatedTime] datetime NOT NULL
) ON [PRIMARY]
GO
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'用户权限ID',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserPermit',@level2type=N'Column',@level2name=N'Id';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'权限代码',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserPermit',@level2type=N'Column',@level2name=N'PermitCode';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'用户ID',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserPermit',@level2type=N'Column',@level2name=N'UserId';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'乐观锁',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserPermit',@level2type=N'Column',@level2name=N'Revision';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserPermit',@level2type=N'Column',@level2name=N'CreatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserPermit',@level2type=N'Column',@level2name=N'CreatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserPermit',@level2type=N'Column',@level2name=N'UpdatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserPermit',@level2type=N'Column',@level2name=N'UpdatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'用户直接权限',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserPermit';
GO

CREATE TABLE [dbo].[AuthUserProperty] (
	[Id] varchar(32) NOT NULL PRIMARY KEY CLUSTERED,
	[UserId] varchar(32) NOT NULL,
	[Name] nvarchar(48) NOT NULL,
	[Value] nvarchar(1024) NOT NULL,
	[Revision] int NOT NULL,
	[CreatedBy] varchar(32) NOT NULL,
	[CreatedTime] datetime NOT NULL,
	[UpdatedBy] varchar(32) NOT NULL,
	[UpdatedTime] datetime NOT NULL
) ON [PRIMARY]
GO
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'属性ID',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserProperty',@level2type=N'Column',@level2name=N'Id';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'用户ID',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserProperty',@level2type=N'Column',@level2name=N'UserId';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'属性名',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserProperty',@level2type=N'Column',@level2name=N'Name';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'属性值',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserProperty',@level2type=N'Column',@level2name=N'Value';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'乐观锁',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserProperty',@level2type=N'Column',@level2name=N'Revision';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserProperty',@level2type=N'Column',@level2name=N'CreatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserProperty',@level2type=N'Column',@level2name=N'CreatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserProperty',@level2type=N'Column',@level2name=N'UpdatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserProperty',@level2type=N'Column',@level2name=N'UpdatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'用户属性',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserProperty';
GO

CREATE TABLE [dbo].[AuthUserRole] (
	[Id] varchar(32) NOT NULL PRIMARY KEY CLUSTERED,
	[OrgId] varchar(32) NOT NULL,
	[RoleId] varchar(32) NOT NULL,
	[UserId] varchar(32) NOT NULL,
	[PositionType] varchar(32) NOT NULL,
	[Revision] int NOT NULL,
	[CreatedBy] varchar(32) NOT NULL,
	[CreatedTime] datetime NOT NULL,
	[UpdatedBy] varchar(32) NOT NULL,
	[UpdatedTime] datetime NOT NULL
) ON [PRIMARY]
GO
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'职责ID',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserRole',@level2type=N'Column',@level2name=N'Id';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'机构ID',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserRole',@level2type=N'Column',@level2name=N'OrgId';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'角色ID',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserRole',@level2type=N'Column',@level2name=N'RoleId';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'用户ID',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserRole',@level2type=N'Column',@level2name=N'UserId';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'岗位类型',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserRole',@level2type=N'Column',@level2name=N'PositionType';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'乐观锁',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserRole',@level2type=N'Column',@level2name=N'Revision';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserRole',@level2type=N'Column',@level2name=N'CreatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'创建时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserRole',@level2type=N'Column',@level2name=N'CreatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新人',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserRole',@level2type=N'Column',@level2name=N'UpdatedBy';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'更新时间',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserRole',@level2type=N'Column',@level2name=N'UpdatedTime';
EXEC sp_addextendedproperty @name=N'MS_Description',@value=N'用户角色',@level0type=N'Schema',@level0name=N'dbo',@level1type=N'Table',@level1name=N'AuthUserRole';
GO