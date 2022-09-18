create table categoria
(
    id_categoria int auto_increment
        primary key,
    nombre       varchar(50) not null
);

create table pais
(
    id_pais int auto_increment
        primary key,
    nombre  varchar(75) not null
);

create table cliente
(
    id_cliente int auto_increment
        primary key,
    nombre     varchar(75)    not null,
    apellido   varchar(75)    not null,
    direccion  varchar(100)   not null,
    telefono   bigint         not null,
    tarjeta    bigint         not null,
    edad       smallint       not null,
    salario    decimal(10, 2) not null,
    genero     char           not null,
    id_pais    int            not null,
    constraint cliente_pais_null_fk
        foreign key (id_pais) references pais (id_pais)
);

create table orden
(
    id_orden   int auto_increment
        primary key,
    fecha      date not null,
    id_cliente int  not null,
    constraint orden_cliente_null_fk
        foreign key (id_cliente) references cliente (id_cliente)
);

create table producto
(
    id_producto  int auto_increment
        primary key,
    nombre       varchar(75)    not null,
    precio       decimal(10, 2) null,
    id_categoria int            not null,
    constraint producto_categoria_null_fk
        foreign key (id_categoria) references categoria (id_categoria)
);

create table temp_orden
(
    id_orden    int  null,
    linea_orden int  null,
    fecha_orden date null,
    id_cliente  int  null,
    id_vendedor int  null,
    id_producto int  null,
    cantidad    int  null
);

create table vendedor
(
    id_vendedor int auto_increment
        primary key,
    nombre      varchar(75) null,
    id_pais     int         not null,
    constraint vendedor_pais_null_fk
        foreign key (id_pais) references pais (id_pais)
);

create table detorden
(
    id_detorden int auto_increment
        primary key,
    cantidad    int not null,
    id_producto int not null,
    id_orden    int not null,
    id_vendedor int not null,
    constraint detorden_orden_null_fk
        foreign key (id_orden) references orden (id_orden),
    constraint detorden_producto_null_fk
        foreign key (id_producto) references producto (id_producto),
    constraint detorden_vendedor_null_fk
        foreign key (id_vendedor) references vendedor (id_vendedor)
);

