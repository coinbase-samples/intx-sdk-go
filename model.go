/**
 * Copyright 2024-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package intx

type ErrorMessage struct {
	Value string `json:"message"`
}

type Portfolio struct {
	PortfolioId    string `json:"portfolio_id"`
	PortfolioUuid  string `json:"portfolio_uuid"`
	Name           string `json:"name"`
	UserUuid       string `json:"user_uuid"`
	MakerFeeRate   string `json:"maker_fee_rate"`
	TakerFeeRate   string `json:"taker_fee_rate"`
	TradingLock    bool   `json:"trading_lock"`
	BorrowDisabled bool   `json:"borrow_disabled"`
	IsLsp          bool   `json:"is_lsp"`
	IsDefault      bool   `json:"is_default"`
}

type Asset struct {
	AssetId                  string  `json:"asset_id"`
	AssetUuid                string  `json:"asset_uuid"`
	AssetName                string  `json:"asset_name"`
	Status                   string  `json:"status"`
	CollateralWeight         float64 `json:"collateral_weight"`
	SupportedNetworksEnabled bool    `json:"supported_networks_enabled"`
}

type Network struct {
	AssetId          string `json:"asset_id"`
	AssetUuid        string `json:"asset_uuid"`
	AssetName        string `json:"asset_name"`
	IsDefault        bool   `json:"is_default"`
	NetworkName      string `json:"network_name"`
	DisplayName      string `json:"display_name"`
	NetworkArnId     string `json:"network_arn_id"`
	MinWithdrawalAmt string `json:"min_withdrawal_amt"`
	MaxWithdrawalAmt string `json:"max_withdrawal_amt"`
	NetworkConfirms  int    `json:"network_confirms"`
	ProcessingTime   int    `json:"processing_time"`
}

type Instrument struct {
	InstrumentId       string  `json:"instrument_id"`
	InstrumentUuid     string  `json:"instrument_uuid"`
	Symbol             string  `json:"symbol"`
	Type               string  `json:"type"`
	BaseAssetId        string  `json:"base_asset_id"`
	BaseAssetUuid      string  `json:"base_asset_uuid"`
	BaseAssetName      string  `json:"base_asset_name"`
	QuoteAssetId       string  `json:"quote_asset_id"`
	QuoteAssetUuid     string  `json:"quote_asset_uuid"`
	QuoteAssetName     string  `json:"quote_asset_name"`
	BaseIncrement      string  `json:"base_increment"`
	QuoteIncrement     string  `json:"quote_increment"`
	MarketOrderPercent float64 `json:"market_order_percent"`
	PriceBandPercent   float64 `json:"price_band_percent"`
	Qty24hr            string  `json:"qty_24hr"`
	Notional24hr       string  `json:"notional_24hr"`
	AvgDailyQty        string  `json:"avg_daily_qty"`
	AvgDailyNotional   string  `json:"avg_daily_notional"`
	PreviousDayQty     string  `json:"previous_day_qty"`
	PositionLimitQty   string  `json:"position_limit_qty"`
	PositionLimitAdv   float64 `json:"position_limit_adv"`
	InitialMarginAdv   float64 `json:"initial_margin_adv"`
	ReplacementCost    string  `json:"replacement_cost"`
	BaseImf            float64 `json:"base_imf"`
	MinNotionalValue   string  `json:"min_notional_value"`
	FundingInterval    string  `json:"funding_interval"`
	TradingState       string  `json:"trading_state"`
	OpenInterest       string  `json:"open_interest"`
	Quote              Quote   `json:"quote"`
}

type Quote struct {
	BestBidPrice     string `json:"best_bid_price"`
	BestBidSize      string `json:"best_bid_size"`
	BestAskPrice     string `json:"best_ask_price"`
	BestAskSize      string `json:"best_ask_size"`
	TradePrice       string `json:"trade_price"`
	TradeQty         string `json:"trade_qty"`
	IndexPrice       string `json:"index_price"`
	MarkPrice        string `json:"mark_price"`
	SettlementPrice  string `json:"settlement_price"`
	LimitUp          string `json:"limit_up"`
	LimitDown        string `json:"limit_down"`
	PredictedFunding string `json:"predicted_funding"`
	Timestamp        string `json:"timestamp"`
}

type HistoricalFundingRate struct {
	InstrumentId string  `json:"instrument_id"`
	FundingRate  float64 `json:"funding_rate"`
	MarkPrice    float64 `json:"mark_price"`
	EventTime    string  `json:"event_time"`
}

type Summary struct {
	Collateral                         string  `json:"collateral"`
	UnrealizedPnl                      string  `json:"unrealized_pnl"`
	PositionNotional                   string  `json:"position_notional"`
	OpenPositionNotional               string  `json:"open_position_notional"`
	PendingFees                        string  `json:"pending_fees"`
	Borrow                             string  `json:"borrow"`
	AccruedInterest                    string  `json:"accrued_interest"`
	RollingDebt                        string  `json:"rolling_debt"`
	Balance                            string  `json:"balance"`
	BuyingPower                        string  `json:"buying_power"`
	PortfolioCurrentMargin             float64 `json:"portfolio_current_margin"`
	PortfolioInitialMargin             float64 `json:"portfolio_initial_margin"`
	PortfolioMaintenanceMargin         float64 `json:"portfolio_maintenance_margin"`
	PortfolioCloseOutMargin            float64 `json:"portfolio_close_out_margin"`
	InLiquidation                      bool    `json:"in_liquidation"`
	PortfolioCurrentMarginNotional     float64 `json:"portfolio_current_margin_notional"`
	PortfolioInitialMarginNotional     float64 `json:"portfolio_initial_margin_notional"`
	PortfolioMaintenanceMarginNotional float64 `json:"portfolio_maintenance_margin_notional"`
	PortfolioCloseOutMarginNotional    float64 `json:"portfolio_close_out_margin_notional"`
	MarginOverride                     float64 `json:"margin_override"`
	LockupInitialMargin                float64 `json:"lockup_initial_margin"`
}

type Balance struct {
	AssetId           string `json:"asset_id"`
	AssetUuid         string `json:"asset_uuid"`
	AssetName         string `json:"asset_name"`
	Quantity          string `json:"quantity"`
	Hold              string `json:"hold"`
	TransferHold      string `json:"transfer_hold"`
	CollateralValue   string `json:"collateral_value"`
	MaxWithdrawAmount string `json:"max_withdraw_amount"`
}

type Position struct {
	InstrumentId   string `json:"instrument_id"`
	InstrumentUuid string `json:"instrument_uuid"`
	Symbol         string `json:"symbol"`
	Vwap           string `json:"vwap"`
	NetSize        string `json:"net_size"`
	BuyOrderSize   string `json:"buy_order_size"`
	SellOrderSize  string `json:"sell_order_size"`
	ImContribution string `json:"im_contribution"`
	UnrealizedPnl  string `json:"unrealized_pnl"`
	MarkPrice      string `json:"mark_price"`
}

type Details struct {
	Summary   Summary    `json:"summary"`
	Balances  []Balance  `json:"balances"`
	Positions []Position `json:"positions"`
}

type Order struct {
	OrderId        string `json:"order_id"`
	ClientOrderId  string `json:"client_order_id"`
	Side           string `json:"side"`
	InstrumentId   string `json:"instrument_id"`
	InstrumentUuid string `json:"instrument_uuid"`
	Symbol         string `json:"symbol"`
	PortfolioId    string `json:"portfolio_id"`
	PortfolioUuid  string `json:"portfolio_uuid"`
	Type           string `json:"type"`
	Price          string `json:"price"`
	StopPrice      string `json:"stop_price"`
	Size           string `json:"size"`
	Tif            string `json:"tif"`
	ExpireTime     string `json:"expire_time"`
	StpMode        string `json:"stp_mode"`
	EventType      string `json:"event_type"`
	OrderStatus    string `json:"order_status"`
	LeavesQty      string `json:"leaves_qty"`
	ExecQty        string `json:"exec_qty"`
	AvgPrice       string `json:"avg_price"`
	Message        string `json:"message"`
	Fee            string `json:"fee"`
}

type MarginOverride struct {
	PortfolioId    string `json:"portfolio_id"`
	MarginOverride string `json:"margin_override"`
}

type Fill struct {
	PortfolioId    string `json:"portfolio_id"`
	PortfolioUuid  string `json:"portfolio_uuid"`
	PortfolioName  string `json:"portfolio_name"`
	FillId         string `json:"fill_id"`
	ExecId         string `json:"exec_id"`
	OrderId        string `json:"order_id"`
	InstrumentId   string `json:"instrument_id"`
	InstrumentUuid string `json:"instrument_uuid"`
	Symbol         string `json:"symbol"`
	MatchId        string `json:"match_id"`
	FillPrice      string `json:"fill_price"`
	FillQty        string `json:"fill_qty"`
	ClientId       string `json:"client_id"`
	ClientOrderId  string `json:"client_order_id"`
	OrderQty       string `json:"order_qty"`
	LimitPrice     string `json:"limit_price"`
	TotalFilled    string `json:"total_filled"`
	FilledVwap     string `json:"filled_vwap"`
	ExpireTime     string `json:"expire_time"`
	StopPrice      string `json:"stop_price"`
	Side           string `json:"side"`
	Tif            string `json:"tif"`
	StpMode        string `json:"stp_mode"`
	Flags          string `json:"flags"`
	Fee            string `json:"fee"`
	FeeAsset       string `json:"fee_asset"`
	OrderStatus    string `json:"order_status"`
	EventTime      string `json:"event_time"`
}

type PaginationSubset struct {
	ResultLimit  int `json:"result_limit"`
	ResultOffset int `json:"result_offset"`
}

type PortfolioSubset struct {
	ID   string `json:"id"`
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type Transfer struct {
	TransferUUID       string          `json:"transfer_uuid"`
	Type               string          `json:"type"`
	Amount             float64         `json:"amount"`
	Asset              string          `json:"asset"`
	Status             string          `json:"status"`
	NetworkName        string          `json:"network_name"`
	CreatedAt          string          `json:"created_at"`
	UpdatedAt          string          `json:"updated_at"`
	FromPortfolio      PortfolioSubset `json:"from_portfolio"`
	ToPortfolio        PortfolioSubset `json:"to_portfolio"`
	FromAddress        float64         `json:"from_address"`
	ToAddress          float64         `json:"to_address"`
	FromCBAccount      string          `json:"from_cb_account"`
	ToCBAccount        string          `json:"to_cb_account"`
	FromCounterpartyID string          `json:"from_counterparty_id"`
	ToCounterpartyID   string          `json:"to_counterparty_id"`
	InstrumentID       int64           `json:"instrument_id"`
	PositionID         string          `json:"position_id"`
}

type Address struct {
	Address      string `json:"address"`
	NetworkArnId string `json:"network_arn_id"`
}

type Counterparty struct {
	PortfolioUuid  string `json:"portfolio_uuid"`
	CounterpartyId string `json:"counterparty_id"`
}

type CounterpartyWithdrawal struct {
	Idem                 string  `json:"idem"`
	PortfolioUuid        string  `json:"portfolio_uuid"`
	SourceCounterpartyId string  `json:"source_counterparty_id"`
	TargetCounterpartyId string  `json:"target_counterparty_id"`
	Asset                string  `json:"asset"`
	Amount               float64 `json:"amount"`
}

type Validation struct {
	CounterpartyId string `json:"counterparty_id"`
	Valid          bool   `json:"valid"`
}

type TransfersResponse struct {
	Pagination PaginationSubset `json:"pagination"`
	Results    []Transfer       `json:"results"`
}

type PaginationParams struct {
	RefDatetime  string `json:"ref_datetime"`
	ResultLimit  int    `json:"result_limit"`
	ResultOffset int    `json:"result_offset"`
}
