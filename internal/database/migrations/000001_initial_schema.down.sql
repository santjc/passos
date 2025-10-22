-- Rollback Initial Schema Migration
-- Drops all tables, types, and indexes in reverse order

-- ============================================================================
-- DROP INDEXES
-- ============================================================================

-- Mapping indexes
DROP INDEX IF EXISTS idx_mapping_deleted_at;
DROP INDEX IF EXISTS idx_mapping_reseller_reference;
DROP INDEX IF EXISTS idx_mapping_unit_id;
DROP INDEX IF EXISTS idx_mapping_option_id;
DROP INDEX IF EXISTS idx_mapping_product_id;

-- Pricing indexes
DROP INDEX IF EXISTS idx_pricing_deleted_at;
DROP INDEX IF EXISTS idx_pricing_booking_id;
DROP INDEX IF EXISTS idx_pricing_availability_id;
DROP INDEX IF EXISTS idx_pricing_unit_id;
DROP INDEX IF EXISTS idx_pricing_option_id;

-- Offer indexes
DROP INDEX IF EXISTS idx_offer_deleted_at;
DROP INDEX IF EXISTS idx_offer_code;
DROP INDEX IF EXISTS idx_offer_option_id;
DROP INDEX IF EXISTS idx_offer_product_id;

-- Contact indexes
DROP INDEX IF EXISTS idx_contact_deleted_at;
DROP INDEX IF EXISTS idx_contact_email;

-- Order indexes
DROP INDEX IF EXISTS idx_order_deleted_at;
DROP INDEX IF EXISTS idx_order_contact_id;
DROP INDEX IF EXISTS idx_order_status;

-- Booking indexes
DROP INDEX IF EXISTS idx_booking_deleted_at;
DROP INDEX IF EXISTS idx_booking_created_at;
DROP INDEX IF EXISTS idx_booking_contact_id;
DROP INDEX IF EXISTS idx_booking_availability_id;
DROP INDEX IF EXISTS idx_booking_option_id;
DROP INDEX IF EXISTS idx_booking_product_id;
DROP INDEX IF EXISTS idx_booking_status;
DROP INDEX IF EXISTS idx_booking_uuid;

-- Availability indexes
DROP INDEX IF EXISTS idx_availability_deleted_at;
DROP INDEX IF EXISTS idx_availability_end_time;
DROP INDEX IF EXISTS idx_availability_start_time;
DROP INDEX IF EXISTS idx_availability_status;
DROP INDEX IF EXISTS idx_availability_option_id;

-- Unit indexes
DROP INDEX IF EXISTS idx_unit_deleted_at;
DROP INDEX IF EXISTS idx_unit_type;
DROP INDEX IF EXISTS idx_unit_option_id;

-- Option indexes
DROP INDEX IF EXISTS idx_option_deleted_at;
DROP INDEX IF EXISTS idx_option_is_default;
DROP INDEX IF EXISTS idx_option_product_id;

-- Product indexes
DROP INDEX IF EXISTS idx_product_deleted_at;
DROP INDEX IF EXISTS idx_product_reference;
DROP INDEX IF EXISTS idx_product_internal_name;

-- ============================================================================
-- DROP TABLES (in reverse order of dependencies)
-- ============================================================================

-- Mapping tables
DROP TABLE IF EXISTS mapping CASCADE;

-- Pricing tables
DROP TABLE IF EXISTS extra_pricing CASCADE;
DROP TABLE IF EXISTS pricing_tax CASCADE;
DROP TABLE IF EXISTS pricing CASCADE;

-- Booking related tables
DROP TABLE IF EXISTS resource_allocation_seat CASCADE;
DROP TABLE IF EXISTS resource_allocation CASCADE;
DROP TABLE IF EXISTS unit_item CASCADE;
DROP TABLE IF EXISTS booking_notice CASCADE;
DROP TABLE IF EXISTS booking_question_answer CASCADE;
DROP TABLE IF EXISTS booking_delivery_method CASCADE;
DROP TABLE IF EXISTS delivery_option CASCADE;
DROP TABLE IF EXISTS ticket CASCADE;
DROP TABLE IF EXISTS booking_contact CASCADE;
DROP TABLE IF EXISTS cancellation CASCADE;
DROP TABLE IF EXISTS order_booking CASCADE;
DROP TABLE IF EXISTS booking CASCADE;
DROP TABLE IF EXISTS "order" CASCADE;
DROP TABLE IF EXISTS contact_locale CASCADE;
DROP TABLE IF EXISTS contact CASCADE;

-- Availability related tables
DROP TABLE IF EXISTS availability_resource CASCADE;
DROP TABLE IF EXISTS seat CASCADE;
DROP TABLE IF EXISTS resource CASCADE;
DROP TABLE IF EXISTS resource_group CASCADE;
DROP TABLE IF EXISTS availability_offer CASCADE;
DROP TABLE IF EXISTS availability_pickup_point CASCADE;
DROP TABLE IF EXISTS availability_notice CASCADE;
DROP TABLE IF EXISTS opening_hours CASCADE;
DROP TABLE IF EXISTS availability CASCADE;
DROP TABLE IF EXISTS notice CASCADE;
DROP TABLE IF EXISTS tour_group CASCADE;

-- Offer related tables
DROP TABLE IF EXISTS offer_combination_unit CASCADE;
DROP TABLE IF EXISTS offer_combination CASCADE;
DROP TABLE IF EXISTS offer_discount_tax CASCADE;
DROP TABLE IF EXISTS offer_discount CASCADE;
DROP TABLE IF EXISTS tax CASCADE;
DROP TABLE IF EXISTS offer_restrictions_unit CASCADE;
DROP TABLE IF EXISTS offer_restrictions CASCADE;
DROP TABLE IF EXISTS offer CASCADE;

-- Unit related tables
DROP TABLE IF EXISTS unit_contact_field CASCADE;
DROP TABLE IF EXISTS unit_accompanied_by CASCADE;
DROP TABLE IF EXISTS unit CASCADE;
DROP TABLE IF EXISTS unit_content CASCADE;

-- Shared tables
DROP TABLE IF EXISTS option_pickup_point CASCADE;
DROP TABLE IF EXISTS pickup_point CASCADE;
DROP TABLE IF EXISTS meeting_point CASCADE;

-- Option related tables
DROP TABLE IF EXISTS option_contact_field CASCADE;
DROP TABLE IF EXISTS option_availability_time CASCADE;
DROP TABLE IF EXISTS option CASCADE;
DROP TABLE IF EXISTS option_pickups CASCADE;
DROP TABLE IF EXISTS itinerary CASCADE;
DROP TABLE IF EXISTS option_content CASCADE;
DROP TABLE IF EXISTS point CASCADE;
DROP TABLE IF EXISTS point_group CASCADE;

-- Question related tables
DROP TABLE IF EXISTS question_answer CASCADE;
DROP TABLE IF EXISTS question_select_option CASCADE;
DROP TABLE IF EXISTS question CASCADE;

-- Product related tables
DROP TABLE IF EXISTS product_delivery_method CASCADE;
DROP TABLE IF EXISTS product_delivery_format CASCADE;
DROP TABLE IF EXISTS product CASCADE;
DROP TABLE IF EXISTS product_questions CASCADE;
DROP TABLE IF EXISTS product_package CASCADE;
DROP TABLE IF EXISTS product_pricing_currency CASCADE;
DROP TABLE IF EXISTS product_pricing CASCADE;
DROP TABLE IF EXISTS product_content CASCADE;

-- ============================================================================
-- DROP TYPES
-- ============================================================================

DROP TYPE IF EXISTS availability_status CASCADE;
DROP TYPE IF EXISTS order_status CASCADE;
DROP TYPE IF EXISTS booking_status CASCADE;
DROP TYPE IF EXISTS unit_type CASCADE;
DROP TYPE IF EXISTS reseller_status CASCADE;
DROP TYPE IF EXISTS net_discount CASCADE;
DROP TYPE IF EXISTS duration_unit CASCADE;
DROP TYPE IF EXISTS contact_field CASCADE;
DROP TYPE IF EXISTS input_type CASCADE;
DROP TYPE IF EXISTS pricing_per CASCADE;
DROP TYPE IF EXISTS redemption_method CASCADE;
DROP TYPE IF EXISTS delivery_method CASCADE;
DROP TYPE IF EXISTS delivery_format CASCADE;
DROP TYPE IF EXISTS availability_type CASCADE;

