create type stock_indicator_data as
    (
    country varchar(20),
    symbol  varchar(20),
    date    date,
    price   numeric,
    value   numeric
    );

alter type stock_indicator_data owner to root;

create table financial_statement
(
    id                                          bigserial
        constraint "PK_4e795d12a43006ece4788e13371"
            primary key,
    symbol                                      varchar(20) not null,
    date                                        date        not null,
    reported_currency                           varchar(10),
    cik                                         varchar(20),
    filling_date                                date,
    accepted_date                               timestamp,
    calendar_year                               smallint,
    revenue                                     numeric,
    cost_of_revenue                             numeric,
    gross_profit                                numeric,
    gross_profit_ratio                          numeric,
    research_and_development_expenses           numeric,
    general_and_administrative_expenses         numeric,
    selling_and_marketing_expenses              numeric,
    selling_general_and_administrative_expenses numeric,
    other_expenses                              numeric,
    operating_expenses                          numeric,
    cost_and_expenses                           numeric,
    interest_income                             numeric,
    interest_expense                            numeric,
    depreciation_and_amortization               numeric,
    ebitda                                      numeric,
    ebitdaratio                                 numeric,
    operating_income                            numeric,
    operating_income_ratio                      numeric,
    total_other_income_expenses_net             numeric,
    income_before_tax                           numeric,
    income_before_tax_ratio                     numeric,
    income_tax_expense                          numeric,
    net_income                                  numeric,
    net_income_ratio                            numeric,
    eps                                         numeric,
    epsdiluted                                  numeric,
    weighted_average_shs_out                    numeric,
    weighted_average_shs_out_dil                numeric,
    link                                        varchar(255),
    final_link                                  varchar(255),
    period                                      varchar(20)
);

alter table financial_statement
    owner to root;

create unique index financial_statement_symbol_date_idx
    on financial_statement (symbol, date);

create table stock
(
    id                   serial
        constraint "PK_092bc1fc7d860426a1dec5aa8e9"
            primary key,
    symbol               varchar(20)           not null,
    company_name         varchar(200)          not null,
    market_cap           bigint                not null,
    sector               varchar(200)          not null,
    industry             varchar(200)          not null,
    beta                 numeric,
    price                numeric,
    last_annual_dividend numeric               not null,
    volume               numeric               not null,
    exchange             varchar(200)          not null,
    exchange_short_name  varchar(200)          not null,
    country              varchar(200)          not null,
    is_etf               boolean               not null,
    is_actively_trading  boolean               not null,
    has_active_data      boolean default false not null,
    tags                 varchar(255)[],
    constraint symbol_country_uniq
        unique (symbol, country)
);

alter table stock
    owner to root;

create index stock_symbol_idx
    on stock (symbol);

create table news
(
    symbol       varchar(20)   not null,
    publish_date timestamp     not null,
    title        varchar(1024) not null,
    image        varchar(1024),
    site         varchar(64)   not null,
    text         text          not null,
    url          varchar(1024) not null,
    id           bigserial
        constraint "PK_39a43dfcb6007180f04aff2357e"
            primary key
);

alter table news
    owner to root;

create unique index "IDX_103d7635ca323f94ca91c3ce0d"
    on news (symbol, publish_date, title);

create table sector_market_cap
(
    sector     varchar(200) not null,
    industry   varchar(200) not null,
    date       date         not null,
    market_cap bigint,
    primary key (sector, industry, date)
);

alter table sector_market_cap
    owner to root;

create unique index sector_market_cap_sector_industry_date_idx
    on sector_market_cap (sector, industry, date);

create table stock_data
(
    id        serial
        constraint "PK_1366dbd374829c34ab2dc4c661b"
            primary key,
    symbol    varchar(20) not null,
    country   varchar(20) not null,
    date      date        not null,
    low       numeric     not null,
    high      numeric     not null,
    open      numeric     not null,
    close     numeric     not null,
    volume    bigint      not null,
    prev_date date
);

alter table stock_data
    owner to root;

create unique index stock_data_country_symbol_date_idx
    on stock_data (symbol, date);

create index stock_data_prev_date_idx
    on stock_data (prev_date);

create index stock_data_symbol_idx
    on stock_data (symbol);

create table sp500_stock
(
    id               serial
        constraint "PK_9376a20a9e06b8d5df55ffd93a1"
            primary key,
    symbol           varchar(20)  not null,
    name             varchar(200) not null,
    sector           varchar(200) not null,
    sub_sector       varchar(200) not null,
    head_quarter     varchar(200) not null,
    date_first_added date,
    cik              bigint       not null,
    founded          smallint
);

alter table sp500_stock
    owner to root;

create table stock_data_m
(
    id     bigserial
        primary key,
    symbol varchar(20) not null,
    date   timestamp   not null,
    low    numeric     not null,
    high   numeric     not null,
    open   numeric     not null,
    close  numeric     not null,
    volume bigint      not null,
    unique (symbol, date)
);

alter table stock_data_m
    owner to root;

create table stock_ignore
(
    country varchar(20) not null,
    symbol  varchar(20) not null,
    primary key (country, symbol)
);

alter table stock_ignore
    owner to root;

create table trend_index
(
    id     bigserial
        constraint trend_index_pk
            primary key,
    symbol varchar(32),
    key    varchar(32) not null,
    begin  timestamp   not null,
    "end"  timestamp,
    color  varchar(32)
);

alter table trend_index
    owner to root;

create unique index trend_index_begin_end_key_symbol_uindex
    on trend_index (begin, "end", key, symbol);

create table series
(
    id     bigserial
        constraint series_pk
            primary key,
    date   timestamp   not null,
    symbol varchar(20) not null,
    key    varchar(32) not null,
    value  numeric     not null
);

alter table series
    owner to root;

create unique index series_timestamp_symbol_key_uindex
    on series (date, symbol, key);

create table stock_data_marker
(
    id     bigserial
        constraint stock_data_marker_pk
            primary key,
    symbol varchar(32) not null,
    date   date        not null,
    key    varchar(32) not null
);

alter table stock_data_marker
    owner to root;

create unique index stock_data_marker_symbol_date_key_uindex
    on stock_data_marker (symbol, date, key);

create table stock_fake_data
(
    id         bigserial
        primary key,
    symbol     varchar(20) not null,
    country    varchar(20) not null,
    generation varchar(20) not null,
    date       date        not null,
    low        numeric     not null,
    high       numeric     not null,
    open       numeric     not null,
    close      numeric     not null,
    volume     bigint      not null
);

alter table stock_fake_data
    owner to root;

create unique index stock_fake_data_symbol_generation_date_idx
    on stock_fake_data (symbol, generation, date);

create index stock_fake_data_generation_symbol_idx
    on stock_fake_data (generation, symbol);

create table economic_calendar
(
    id                bigserial
        primary key,
    event             varchar(255) not null,
    date              timestamp    not null,
    country           varchar(255),
    actual            numeric,
    previous          numeric,
    change            numeric,
    change_percentage numeric,
    estimate          numeric,
    impact            varchar(10)
);

alter table economic_calendar
    owner to root;

create unique index economic_calendar_event_date_country_idx
    on economic_calendar (event, date, country);

create materialized view market_move_by_year as
SELECT s.sector,
       s.industry,
       year.year,
       round(sum(s.market_cap::numeric *
                             fmp.indicator_avg_trend(s.country, s.symbol, (year.year || '-01-11'::text)::date,
                                                     ((year.year + 1) || '-01-01'::text)::date)))                AS market_cap_move,
       round(avg(fmp.indicator_avg_trend(s.country, s.symbol, (year.year || '-01-11'::text)::date,
                                         ((year.year + 1) || '-01-01'::text)::date)) *
             100::numeric)                                                                           AS market_price_move
FROM fmp.stock s
         JOIN generate_series(2010, 2022) year(year) ON true
WHERE true
  AND s.market_cap > 1000000000
GROUP BY s.sector, s.industry, year.year
ORDER BY year.year,
    (round(sum(s.market_cap::numeric *
    fmp.indicator_avg_trend(s.country, s.symbol, (year.year || '-01-11'::text)::date,
    ((year.year + 1) || '-01-01'::text)::date)))) DESC;

alter materialized view market_move_by_year owner to root;

create materialized view share_count as
SELECT fs.country,
       fs.symbol,
       fs.exchange_short_name,
       round(fs.market_cap::numeric / fs.price) AS share_count
FROM fmp.stock fs
WHERE fs.price > 0::numeric;

alter materialized view share_count owner to root;

create materialized view stock_dates as
SELECT DISTINCT stock_data.date
FROM fmp.stock_data;

alter materialized view stock_dates owner to root;

create view sector_market_cap_view(sector, date, low, high, open, close, volume) as
SELECT s.sector,
       sdx.date,
       sum(sdx.low * sc.share_count)             AS low,
       sum(sdx.high * sc.share_count)            AS high,
       sum(sdx.open * sc.share_count)            AS open,
       sum(sdx.close * sc.share_count)           AS close,
       sum(sdx.volume::numeric * sc.share_count) AS volume
FROM fmp.stock s
         JOIN fmp.stock_data sdx ON s.symbol::text = sdx.symbol::text AND s.country::text = sdx.country::text
         JOIN fmp.share_count sc ON s.symbol::text = sc.symbol::text AND s.country::text = sc.country::text
GROUP BY s.sector, sdx.date;

alter table sector_market_cap_view
    owner to root;

create view trend_index_view(id, symbol, key, begin, "end", open, close) as
SELECT trend_index.id,
       trend_index.symbol,
       trend_index.key,
       trend_index.begin,
       trend_index."end",
       begin_data.open,
       end_data.close
FROM fmp.trend_index
         JOIN fmp.stock_data begin_data
              ON begin_data.symbol::text = trend_index.symbol::text AND begin_data.date = trend_index.begin
         JOIN fmp.stock_data end_data
              ON end_data.symbol::text = trend_index.symbol::text AND end_data.date = trend_index."end";

alter table trend_index_view
    owner to root;

create view stock_data_h(symbol, date, open, close, high, low, volume) as
SELECT stock_data_m.symbol,
       date_trunc('hour'::text, stock_data_m.date)                        AS date,
       (array_agg(stock_data_m.open ORDER BY stock_data_m.date))[1]       AS open,
       (array_agg(stock_data_m.close ORDER BY stock_data_m.date DESC))[1] AS close,
       max(stock_data_m.high)                                             AS high,
       min(stock_data_m.low)                                              AS low,
       sum(stock_data_m.volume)                                           AS volume
FROM fmp.stock_data_m
GROUP BY stock_data_m.symbol, (date_trunc('hour'::text, stock_data_m.date))
UNION
SELECT stock_data_m.symbol,
       date_trunc('hour'::text, stock_data_m.date) + '00:30:00'::interval AS date,
       (array_agg(stock_data_m.open ORDER BY stock_data_m.date))[1]       AS open,
       (array_agg(stock_data_m.close ORDER BY stock_data_m.date DESC))[1] AS close,
       max(stock_data_m.high)                                             AS high,
       min(stock_data_m.low)                                              AS low,
       sum(stock_data_m.volume)                                           AS volume
FROM fmp.stock_data_m
WHERE date_part('hour'::text, stock_data_m.date) = 13::double precision
  AND date_part('minute'::text, stock_data_m.date) >= 30::double precision
GROUP BY stock_data_m.symbol, (date_trunc('hour'::text, stock_data_m.date) + '00:30:00'::interval);

alter table stock_data_h
    owner to root;

create view report_sd_1
            (symbol, date, sd_diff, sd_full_diff, open_gap, sdh_diff, sdh_full_diff, sd_volume, sdh_volume) as
SELECT sd.symbol,
       sd.date,
       sd.close - sd.open   AS sd_diff,
       sd.high - sd.low     AS sd_full_diff,
       sd.open - prev.close AS open_gap,
       sdh.close - sdh.open AS sdh_diff,
       sdh.high - sdh.low   AS sdh_full_diff,
       sd.volume            AS sd_volume,
       sdh.volume           AS sdh_volume
FROM fmp.stock_data sd
         JOIN fmp.stock_data_h sdh ON date_part('hour'::text, sdh.date) = 13::double precision AND
                                      date_part('minute'::text, sdh.date) = 30::double precision AND
                                      sd.symbol::text = sdh.symbol::text AND sd.date = date_trunc('day'::text, sdh.date)
         JOIN fmp.stock_data prev ON prev.date = sd.prev_date AND prev.symbol::text = sd.symbol::text
ORDER BY sd.date DESC;

alter table report_sd_1
    owner to root;

create view report_gap
            (symbol, date, sd_diff, sd_full_diff, open_gap, sd_volume, gap_closed, prev_close, open, close, high,
             low) as
SELECT sd.symbol,
       sd.date,
       sd.close - sd.open                             AS sd_diff,
       sd.high - sd.low                               AS sd_full_diff,
       sd.open - prev.close                           AS open_gap,
       sd.volume                                      AS sd_volume,
       prev.close >= sd.low AND prev.close <= sd.high AS gap_closed,
       prev.close                                     AS prev_close,
       sd.open,
       sd.close,
       sd.high,
       sd.low
FROM fmp.stock_data sd
         JOIN fmp.stock_data prev ON prev.date = sd.prev_date AND prev.symbol::text = sd.symbol::text
WHERE abs(sd.open - prev.close) > 0::numeric
  AND sd.date <> prev.date;

alter table report_gap
    owner to root;

create view report_gap_analyse_1
            (symbol, market_cap, sum_gap, sum_closed_gap, sum_not_closed_gap, count, count_closed_gap, percentage) as
SELECT s.symbol,
       s.market_cap,
       sum(abs(report_gap.open_gap))     AS sum_gap,
       sum(abs(
               CASE
                   WHEN report_gap.gap_closed THEN report_gap.open_gap
                   ELSE 0::numeric
                   END))                 AS sum_closed_gap,
       sum(abs(
               CASE
                   WHEN NOT report_gap.gap_closed THEN
                       CASE
                           WHEN report_gap.open_gap > 0::numeric THEN report_gap.open - report_gap.close
                           ELSE report_gap.close - report_gap.open
                           END
                   ELSE 0::numeric
                   END))                 AS sum_not_closed_gap,
       count(*)                          AS count,
       sum(
               CASE
                   WHEN report_gap.gap_closed THEN 1
                   ELSE 0
                   END)                  AS count_closed_gap,
       sum(
               CASE
                   WHEN report_gap.gap_closed THEN 1
                   ELSE 0
                   END) * 100 / count(*) AS percentage
FROM fmp.report_gap
         JOIN fmp.stock s ON s.symbol::text = report_gap.symbol::text
WHERE s.market_cap > '10000000000'::bigint
  AND report_gap.open_gap > 0.1
  AND report_gap.date > '2022-01-01'::date
GROUP BY s.symbol, s.market_cap;

alter table report_gap_analyse_1
    owner to root;

create view stock_data_variation(symbol, date, variation, high_variation, variation_div) as
SELECT sd.symbol,
       sd.date,
       sum(abs(sdm.close - sdm.open))                           AS variation,
       sum(abs(sdm.high - sdm.low))                             AS high_variation,
       sum(abs(sdm.close - sdm.open)) / abs(sd.close - sd.open) AS variation_div
FROM fmp.stock_data_m sdm
         JOIN fmp.stock_data sd ON sd.symbol::text = sdm.symbol::text AND sd.date = sdm.date::date
GROUP BY sd.id;

alter table stock_data_variation
    owner to root;

create view trend_reversal
            (symbol, date, left_length, right_length, left_diff, right_diff, left_diff_p, right_diff_p) as
SELECT ti1.symbol,
       ti1."end"::date                        AS date,
       ti1."end"::date - ti1.begin::date      AS left_length,
       ti2."end"::date - ti2.begin::date      AS right_length,
       sd1e.close - sd1b.close                AS left_diff,
       sd2e.close - sd2b.close                AS right_diff,
       (sd1e.close - sd1b.close) / sd1e.close AS left_diff_p,
       (sd2e.close - sd2b.close) / sd2e.close AS right_diff_p
FROM fmp.trend_index ti1
         JOIN fmp.trend_index ti2
              ON ti1.key::text = ti2.key::text AND ti1.symbol::text = ti2.symbol::text AND ti1."end"::date = ti2.begin::date
         JOIN fmp.stock_data sd1b ON sd1b.symbol::text = '^NDX'::text AND sd1b.date = ti1.begin::date
         JOIN fmp.stock_data sd1e ON sd1e.symbol::text = '^NDX'::text AND sd1e.date = ti1."end"::date
         JOIN fmp.stock_data sd2b ON sd2b.symbol::text = '^NDX'::text AND sd2b.date = ti2.begin::date
         JOIN fmp.stock_data sd2e ON sd2e.symbol::text = '^NDX'::text AND sd2e.date = ti2."end"::date
WHERE ti1.key::text = 'tdw-trend-arrows'::text;

alter table trend_reversal
    owner to root;

create procedure compute_sector_market_cap(_from_year date, _to_year date)
    language plpgsql
as
$$
DECLARE
list CURSOR FOR select sector, industry
                from fmp.stock s
                where country = 'US'
                  and not exists(select 1
                                 from fmp.stock_ignore as si
                                 where s.country = si.country
                                   and s.symbol = si.symbol)
                group by 1, 2;
dates CURSOR FOR select date
                 from fmp.stock_dates
                 where date between _from_year and _to_year;
exists bool;
count integer = 0;
    checked integer = 0;
BEGIN
for item in list
        LOOP
            raise notice 'Computing: %; %; %; %', item.sector, item.industry, count, checked;
for d in dates
                LOOP
select count(*) > 0
into exists
from fmp.sector_market_cap smc
where smc.sector = item.sector
  and smc.industry = item.industry
  and smc.date = d.date;
checked = checked + 1;
                     if checked % 1000 = 0 then
                            raise notice '%, %', checked, count;
end if;

                    if count > 10000 then
                        raise notice 'finishing on count, %', count;
                        return;
end if;

                    if not exists then
                        count = count + 1;
insert into fmp.sector_market_cap (select item.sector,
                                          item.industry,
                                          d.date,
                                          round(sum(coalesce(
                                                  (fmp.stock_market_cap(s.country, s.symbol, d.date)),
                                                  0))) as market_cap
                                   from fmp.stock s
                                   where s.sector = item.sector
                                     and s.industry = item.industry);
end if;
end loop;
end loop;
end;
$$;

alter procedure compute_sector_market_cap(date, date) owner to root;

create function indicator_avg_trend(_country character varying, _symbol character varying, _from date, _to date) returns numeric
    language plpgsql
as
$$
DECLARE
_left  numeric = 0;
    _right numeric = 0;
    middle date;
BEGIN
    middle = _from + (_to - _from) / 2;

select round(avg(close), 6)
into _left
from fmp.stock_data sdp
where sdp.country = _country
  and sdp.symbol = _symbol
  and sdp.date between _from and middle;

select round(avg(close), 6)
into _right
from fmp.stock_data sdp
where sdp.country = _country
  and sdp.symbol = _symbol
  and sdp.date between middle and _to;

if (_left = 0) then
        return 0;
end if;

return (_right - _left) / _left;
end;
$$;

alter function indicator_avg_trend(varchar, varchar, date, date) owner to root;

create function indicator_mavg(_country character varying, _symbol character varying, _date date, _window integer, _shift integer) returns numeric
    language plpgsql
as
$$
DECLARE
result numeric;
BEGIN
select round(avg(close), 6)
into result
from fmp.stock_data sdp
where sdp.country = _country
  and sdp.symbol = _symbol
  and sdp.date between _date - ((_window - _shift) || ' days')::interval AND _date + (_shift || ' days') :: interval;

return result;
end;
$$;

alter function indicator_mavg(varchar, varchar, date, integer, integer) owner to root;

create function indicator_sector_normalize(_country character varying, _symbol character varying, _date date, _sector character varying) returns numeric
    language plpgsql
as
$$
DECLARE
result numeric;
BEGIN
select smc.close / 30000000000
into result
from fmp.stock_data sd
         inner join fmp.sector_market_cap_view smc on smc.date = sd.date and smc.sector = _sector
where symbol = _symbol
  and country = _country
  and sd.date = _date;

return result;
end;
$$;

alter function indicator_sector_normalize(varchar, varchar, date, varchar) owner to root;

create function stock_market_cap(_country character varying, _symbol character varying, _date date) returns numeric
    language plpgsql
as
$$
declare
result numeric;
BEGIN
select c.close * sc.share_count
into result
from fmp.stock_data c
         inner join fmp.share_count sc on c.symbol = sc.symbol and c.country = sc.country
where c.symbol = _symbol
  and c.date <= _date
  and c.country = _country
order by c.date desc
    limit 1;

return result;
end;
$$;

alter function stock_market_cap(varchar, varchar, date) owner to root;

create function stock_prev_date(_symbol character varying, _before date) returns date
    language plpgsql
as
$$
declare
result date;
BEGIN
select c.date
into result
from fmp.stock_data c
where c.symbol = _symbol
  and c.date < _before
order by c.date desc
    limit 1;

return result;
end;
$$;

alter function stock_prev_date(varchar, date) owner to root;

create function stock_price(_country character varying, _symbol character varying, _before date) returns numeric
    language plpgsql
as
$$
declare
result numeric;
BEGIN
select c.close
into result
from fmp.stock_data c
where c.symbol = _symbol
  and c.date <= _before
  and c.country = _country
order by c.date desc
    limit 1;

return result;
end;
$$;

alter function stock_price(varchar, varchar, date) owner to root;

create function fix_prev_date() returns trigger
    language plpgsql
as
$$
BEGIN
update fmp.stock_data
set prev_date = fmp.stock_prev_date(stock_data.symbol, stock_data.date)
where id = NEW.id;

return NEW;
END;
$$;

alter function fix_prev_date() owner to root;

create trigger stock_data_prev_date_fix
    after insert
    on stock_data
    execute procedure fix_prev_date();

