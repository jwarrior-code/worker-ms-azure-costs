CREATE TABLE azure_prices (
                              skuId VARCHAR(250) PRIMARY KEY,
                              currencyCode VARCHAR(50) NOT NULL,
                              tierMinimumUnits VARCHAR(255),
                              retailPrice VARCHAR(255),
                              armRegionName VARCHAR(100),
                              location VARCHAR(100),
                              effectiveStartDate VARCHAR(50),
                              productName VARCHAR(255),
                              skuName VARCHAR(255),
                              serviceName VARCHAR(255),
                              armSkuName VARCHAR(255)
);