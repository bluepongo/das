package metadata

import (
	"fmt"

	"github.com/pingcap/errors"
	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/log"

	"github.com/romberli/das/global"
)

var _ metadata.MiddlewareClusterRepo = (*MiddlewareClusterRepo)(nil)

type MiddlewareClusterRepo struct {
	Database middleware.Pool
}

// NewMiddlewareClusterRepo returns *MiddlewareClusterRepo with given middleware.Pool
func NewMiddlewareClusterRepo(db middleware.Pool) *MiddlewareClusterRepo {
	return &MiddlewareClusterRepo{db}
}

// NewMiddlewareClusterRepoWithGlobal returns *MiddlewareClusterRepo with global mysql pool
func NewMiddlewareClusterRepoWithGlobal() *MiddlewareClusterRepo {
	return NewMiddlewareClusterRepo(global.DASMySQLPool)
}

// Execute executes command with arguments on database
func (mcr *MiddlewareClusterRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := mcr.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata MiddlewareClusterRepo.Execute(): close database connection failed.\n%+v", err)
		}
	}()

	return conn.Execute(command, args...)
}

// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
func (mcr *MiddlewareClusterRepo) Transaction() (middleware.Transaction, error) {
	return mcr.Database.Transaction()
}

// GetAll gets all middleware clusters from the middleware
func (mcr *MiddlewareClusterRepo) GetAll() ([]metadata.MiddlewareCluster, error) {
	sql := `
		select id, cluster_name, env_id, del_flag, create_time, last_update_time
		from t_meta_middleware_cluster_info
		where del_flag = 0
		order by id;
	`
	log.Debugf("metadata MiddlewareClusterRepo.GetAll() sql: \n%s", sql)

	result, err := mcr.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*MiddlewareClusterInfo
	middlewareClusterInfoList := make([]*MiddlewareClusterInfo, result.RowNumber())
	for i := range middlewareClusterInfoList {
		middlewareClusterInfoList[i] = NewEmptyMiddlewareClusterInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(middlewareClusterInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []dependency.Entity
	entityList := make([]metadata.MiddlewareCluster, result.RowNumber())
	for i := range entityList {
		entityList[i] = middlewareClusterInfoList[i]
	}

	return entityList, nil
}

// GetByEnv gets middleware clusters of given env id from the middleware
func (mcr *MiddlewareClusterRepo) GetByEnv(envID int) ([]metadata.MiddlewareCluster, error) {
	sql := `
		select id, cluster_name, env_id, del_flag, create_time, last_update_time
		from t_meta_middleware_cluster_info
		where del_flag = 0
		and env_id = ?
		order by id;
	`
	log.Debugf("metadata MiddlewareClusterRepo.GetByEnv() sql: \n%s", sql, envID)

	result, err := mcr.Execute(sql, envID)
	if err != nil {
		return nil, err
	}
	// init []*MiddlewareClusterInfo
	middlewareClusterInfoList := make([]*MiddlewareClusterInfo, result.RowNumber())
	for i := range middlewareClusterInfoList {
		middlewareClusterInfoList[i] = NewEmptyMiddlewareClusterInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(middlewareClusterInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []metadata.MiddlewareCluster
	middlewareClusterList := make([]metadata.MiddlewareCluster, result.RowNumber())
	for i := range middlewareClusterList {
		middlewareClusterList[i] = middlewareClusterInfoList[i]
	}

	return middlewareClusterList, nil
}

// GetByID gets a middleware cluster by the identity from the middleware
func (mcr *MiddlewareClusterRepo) GetByID(id int) (metadata.MiddlewareCluster, error) {
	sql := `
		select id, cluster_name, env_id, del_flag, create_time, last_update_time
		from t_meta_middleware_cluster_info
		where del_flag = 0
		and id = ?;
	`
	log.Debugf("metadata MiddlewareClusterRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, id)

	result, err := mcr.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata MiddlewareClusterInfo.GetByID(): data does not exists, id: %d", id))
	case 1:
		middlewareClusterInfo := NewEmptyMiddlewareClusterInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(middlewareClusterInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return middlewareClusterInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata MiddlewareClusterInfo.GetByID(): duplicate key exists, id: %d", id))
	}
}

// GetByName gets a middleware cluster of given cluster name from the middle ware
func (mcr *MiddlewareClusterRepo) GetByName(clusterName string) (metadata.MiddlewareCluster, error) {
	sql := `select id from t_meta_middleware_cluster_info where del_flag = 0 and cluster_name = ?;`
	log.Debugf("metadata MiddlewareClusterRepo.GetByName() select sql: %s", sql)
	result, err := mcr.Execute(sql, clusterName)
	if err != nil {
		return nil, err
	}
	id, err := result.GetInt(constant.ZeroInt, constant.ZeroInt)
	if err != nil {
		return nil, err
	}
	return mcr.GetByID(id)
}

// GetID gets the identity with given cluster name and env id from the middleware
func (mcr *MiddlewareClusterRepo) GetID(clusterName string, envID int) (int, error) {
	sql := `select id from t_meta_middleware_cluster_info where del_flag = 0 and cluster_name = ? and env_id = ?;`
	log.Debugf("metadata MiddlewareClusterRepo.GetID() select sql: %s", sql)
	result, err := mcr.Execute(sql, clusterName, envID)
	if err != nil {
		return constant.ZeroInt, err
	}

	return result.GetInt(constant.ZeroInt, constant.ZeroInt)
}

// GetUsersByMiddlewareClusterID get user list that own the middleware cluster
func (mcr *MiddlewareClusterRepo) GetUsersByMiddlewareClusterID(clusterID int) ([]metadata.User, error) {
	sql := `
		select ui.id,
			   ui.user_name,
			   ui.department_name,
			   ui.employee_id,
			   ui.account_name,
			   ui.email,
			   ui.telephone,
			   ui.mobile,
			   ui.role,
			   ui.del_flag,
			   ui.create_time,
			   ui.last_update_time
		from t_meta_middleware_cluster_info mci
				 inner join t_meta_middleware_cluster_user_map mcum
							on mci.id = mcum.middleware_cluster_id
				 inner join t_meta_user_info ui on mcum.user_id = ui.id
		where mci.del_flag = 0
		  and mcum.del_flag = 0
		  and ui.del_flag = 0
		  and mci.id = ?
	`
	log.Debugf("metadata MiddlewareClusterRepo.GetUsersByMiddlewareClusterID() select sql: %s", sql)
	result, err := mcr.Execute(sql, clusterID)
	if err != nil {
		return nil, err
	}

	// init []dependency.Entity
	userList := make([]metadata.User, result.RowNumber())
	for i := range userList {
		userList[i] = NewEmptyUserInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(userList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return userList, nil

}

// Create creates data with given entity in the middleware
func (mcr *MiddlewareClusterRepo) Create(middlewareCluster metadata.MiddlewareCluster) (metadata.MiddlewareCluster, error) {
	sql := `insert into t_meta_middleware_cluster_info(cluster_name, env_id) values(?, ?);`
	log.Debugf("metadata MiddlewareClusterRepo.Create() insert sql: %s", sql)
	// execute
	_, err := mcr.Execute(sql,
		middlewareCluster.(*MiddlewareClusterInfo).ClusterName,
		middlewareCluster.(*MiddlewareClusterInfo).EnvID,
	)
	if err != nil {
		return nil, err
	}
	// get id
	id, err := mcr.GetID(middlewareCluster.GetClusterName(), middlewareCluster.GetEnvID())
	if err != nil {
		return nil, err
	}
	// get entity
	return mcr.GetByID(id)
}

// Update updates data with given entity in the middleware
func (mcr *MiddlewareClusterRepo) Update(middlewareCluster metadata.MiddlewareCluster) error {
	sql := `update t_meta_middleware_cluster_info set cluster_name = ?, env_id = ?, del_flag = ? where id = ?;`
	log.Debugf("metadata MiddlewareClusterRepo.Update() update sql: %s", sql)
	_, err := mcr.Execute(sql,
		middlewareCluster.GetClusterName(),
		middlewareCluster.GetEnvID(),
		middlewareCluster.GetDelFlag(),
		middlewareCluster.Identity(),
	)

	return err
}

// Delete deletes the middleware cluster in the middleware
func (mcr *MiddlewareClusterRepo) Delete(id int) error {
	tx, err := mcr.Transaction()
	if err != nil {
		return err
	}
	defer func() {
		err = tx.Close()
		if err != nil {
			log.Errorf("metadata MiddlewareClusterRepo.Delete(): close database connection failed.\n%+v", err)
		}
	}()

	err = tx.Begin()
	if err != nil {
		return err
	}
	sql := `delete from t_meta_middleware_cluster_info where id = ?;`
	log.Debugf("metadata MiddlewareClusterRepo.Delete() update sql: %s", sql)
	_, err = mcr.Execute(sql, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// AddUser adds a new map of middleware cluster and user in the middleware
func (mcr *MiddlewareClusterRepo) AddUser(middlewareClusterID, userID int) error {
	userRepo := NewUserRepoWithGlobal()
	_, err := userRepo.GetByID(userID)
	if err != nil {
		return nil
	}
	_, err = mcr.GetByID(middlewareClusterID)
	if err != nil {
		return nil
	}
	sql := `insert into t_meta_middleware_cluster_user_map(middleware_cluster_id, user_id) values(?, ?);`
	log.Debugf("metadata MiddlewareClusterRepo.MiddlewareClusterAddUser() insert sql: %s", sql)
	_, err = mcr.Execute(sql, middlewareClusterID, userID)

	return err
}

// DeleteUser deletes a map of middleware cluster and user in the middleware
func (mcr *MiddlewareClusterRepo) DeleteUser(middlewareClusterID, userID int) error {
	userRepo := NewUserRepoWithGlobal()
	_, err := userRepo.GetByID(userID)
	if err != nil {
		return err
	}
	_, err = mcr.GetByID(middlewareClusterID)
	if err != nil {
		return err
	}
	sql := `delete from t_meta_middleware_cluster_user_map where middleware_cluster_id = ? and user_id = ?;`
	log.Debugf("metadata MiddlewareClusterRepo.MiddlewareClusterAddUser() delete sql: %s", sql)
	_, err = mcr.Execute(sql, middlewareClusterID, userID)

	return err
}
