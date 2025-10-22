-- Initial Schema Migration
-- Creates all tables, types, and indexes for the Passos booking system

-- Create ENUM types
CREATE TYPE availability_type AS ENUM ('START_TIME', 'OPENING_HOURS');
CREATE TYPE delivery_format AS ENUM ('PDF_URL', 'QRCODE', 'CODE128', 'PKPASS_URL', 'AZTECCODE', 'GOOGLE_WALLET_URL');
CREATE TYPE delivery_method AS ENUM ('VOUCHER', 'TICKET');
CREATE TYPE redemption_method AS ENUM ('DIGITAL', 'PRINT', 'MANIFEST');
CREATE TYPE pricing_per AS ENUM ('UNIT', 'BOOKING');
CREATE TYPE input_type AS ENUM ('date', 'datetime-local', 'email', 'file', 'number', 'tel', 'text', 'time', 'url', 'select', 'textarea');
CREATE TYPE contact_field AS ENUM ('firstName', 'lastName', 'emailAddress', 'phoneNumber', 'country', 'notes', 'locales', 'allowMarketing', 'postalCode');
CREATE TYPE duration_unit AS ENUM ('hour', 'minute', 'day');
CREATE TYPE net_discount AS ENUM ('NONE', 'FULL', 'SPLIT', 'PRORATED');
CREATE TYPE reseller_status AS ENUM ('ACTIVE', 'DISABLED', 'DRAFT');
CREATE TYPE unit_type AS ENUM ('ADULT', 'YOUTH', 'CHILD', 'INFANT', 'FAMILY', 'SENIOR', 'STUDENT', 'MILITARY', 'OTHER');
CREATE TYPE booking_status AS ENUM ('ON_HOLD', 'CONFIRMED', 'EXPIRED', 'CANCELLED', 'REDEEMED', 'PENDING', 'REJECTED');
CREATE TYPE order_status AS ENUM ('ON_HOLD', 'CONFIRMED', 'CANCELLED', 'EXPIRED', 'PENDING');
CREATE TYPE availability_status AS ENUM ('AVAILABLE', 'FREESALE', 'SOLD_OUT', 'LIMITED', 'CLOSED');

-- ============================================================================
-- PRODUCT RELATED TABLES
-- ============================================================================

-- Product Content
CREATE TABLE product_content (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT,
    country TEXT,
    location TEXT,
    subtitle TEXT,
    short_description TEXT,
    description TEXT,
    highlights TEXT,
    inclusions TEXT,
    exclusions TEXT,
    booking_terms TEXT,
    redemption_instructions TEXT,
    cancellation_policy TEXT,
    destination TEXT,
    categories TEXT,
    faqs TEXT,
    cover_image_url TEXT,
    banner_image_url TEXT,
    video_url TEXT,
    gallery_images TEXT,
    banner_images TEXT,
    point_to_point BOOLEAN,
    privacy_terms TEXT,
    alert TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Product Pricing
CREATE TABLE product_pricing (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    default_currency TEXT,
    pricing_per pricing_per,
    include_tax BOOLEAN,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Product Pricing Currency
CREATE TABLE product_pricing_currency (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_pricing_id UUID NOT NULL REFERENCES product_pricing(id) ON DELETE CASCADE,
    currency TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Product Package
CREATE TABLE product_package (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    is_package BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Product Questions
CREATE TABLE product_questions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Product
CREATE TABLE product (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    internal_name TEXT NOT NULL,
    reference TEXT,
    locale TEXT NOT NULL,
    time_zone TEXT NOT NULL,
    allow_freesale BOOLEAN NOT NULL DEFAULT false,
    instant_confirmation BOOLEAN NOT NULL DEFAULT false,
    instant_delivery BOOLEAN NOT NULL DEFAULT false,
    availability_required BOOLEAN NOT NULL DEFAULT false,
    availability_type availability_type NOT NULL,
    redemption_method redemption_method NOT NULL,
    freesale_duration_amount INTEGER NOT NULL,
    freesale_duration_unit TEXT NOT NULL,
    product_content_id UUID NOT NULL REFERENCES product_content(id) ON DELETE RESTRICT,
    product_pricing_id UUID NOT NULL REFERENCES product_pricing(id) ON DELETE RESTRICT,
    product_package_id UUID NOT NULL REFERENCES product_package(id) ON DELETE RESTRICT,
    product_questions_id UUID REFERENCES product_questions(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Product Delivery Format
CREATE TABLE product_delivery_format (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_id UUID NOT NULL REFERENCES product(id) ON DELETE CASCADE,
    delivery_format delivery_format NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Product Delivery Method
CREATE TABLE product_delivery_method (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_id UUID NOT NULL REFERENCES product(id) ON DELETE CASCADE,
    delivery_method delivery_method NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================================
-- QUESTION RELATED TABLES
-- ============================================================================

-- Question
CREATE TABLE question (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_questions_id UUID NOT NULL REFERENCES product_questions(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    short_description TEXT NOT NULL,
    required BOOLEAN NOT NULL DEFAULT false,
    input_type input_type NOT NULL,
    cover_image_url TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Question Select Option
CREATE TABLE question_select_option (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    question_id UUID NOT NULL REFERENCES question(id) ON DELETE CASCADE,
    label TEXT NOT NULL,
    value TEXT NOT NULL,
    sort_order INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Question Answer
CREATE TABLE question_answer (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    question_id UUID NOT NULL REFERENCES question(id) ON DELETE CASCADE,
    booking_id UUID NOT NULL,
    value TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================================
-- OPTION RELATED TABLES
-- ============================================================================

-- Point Group
CREATE TABLE point_group (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    short_description TEXT NOT NULL,
    internal_name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Point
CREATE TABLE point (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    internal_name TEXT NOT NULL,
    short_description TEXT NOT NULL,
    title TEXT NOT NULL,
    point_group_id UUID REFERENCES point_group(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Option Content
CREATE TABLE option_content (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT,
    subtitle TEXT,
    language TEXT,
    short_description TEXT,
    duration TEXT,
    duration_amount TEXT,
    duration_unit duration_unit,
    cover_image_url TEXT,
    from_point_id UUID REFERENCES point(id) ON DELETE SET NULL,
    to_point_id UUID REFERENCES point(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Itinerary
CREATE TABLE itinerary (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    option_content_id UUID NOT NULL REFERENCES option_content(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    description TEXT NOT NULL,
    address TEXT NOT NULL,
    google_place_id TEXT NOT NULL,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,
    travel_time TEXT NOT NULL,
    travel_time_amount INTEGER NOT NULL,
    travel_time_unit TEXT NOT NULL,
    duration TEXT NOT NULL,
    duration_amount INTEGER NOT NULL,
    duration_unit TEXT NOT NULL,
    sort_order INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Option Pickups
CREATE TABLE option_pickups (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    pickup_required BOOLEAN,
    pickup_available BOOLEAN,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Option
CREATE TABLE option (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_id UUID NOT NULL REFERENCES product(id) ON DELETE CASCADE,
    is_default BOOLEAN NOT NULL DEFAULT false,
    internal_name TEXT NOT NULL,
    reference TEXT,
    cancellation_cutoff TEXT NOT NULL,
    cancellation_cutoff_amount INTEGER NOT NULL,
    cancellation_cutoff_unit TEXT NOT NULL,
    availability_cutoff TEXT NOT NULL,
    availability_cutoff_amount INTEGER NOT NULL,
    availability_cutoff_unit TEXT NOT NULL,
    min_units INTEGER NOT NULL,
    max_units INTEGER NOT NULL,
    min_pax_count INTEGER NOT NULL,
    max_pax_count INTEGER NOT NULL,
    option_content_id UUID NOT NULL REFERENCES option_content(id) ON DELETE RESTRICT,
    option_pickups_id UUID REFERENCES option_pickups(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Option Availability Time
CREATE TABLE option_availability_time (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    option_id UUID NOT NULL REFERENCES option(id) ON DELETE CASCADE,
    local_time TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Option Contact Field
CREATE TABLE option_contact_field (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    option_id UUID NOT NULL REFERENCES option(id) ON DELETE CASCADE,
    field contact_field NOT NULL,
    is_required BOOLEAN NOT NULL DEFAULT false,
    is_visible BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================================
-- SHARED TABLES (Meeting/Pickup Points)
-- ============================================================================

-- Meeting Point
CREATE TABLE meeting_point (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    address TEXT,
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Pickup Point
CREATE TABLE pickup_point (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    directions TEXT,
    address TEXT NOT NULL,
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    google_place_id TEXT,
    street TEXT,
    postal_code TEXT,
    locality TEXT,
    region TEXT,
    state TEXT,
    country TEXT,
    local_date_time TEXT NOT NULL,
    local_date_time_to TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Option Pickup Point
CREATE TABLE option_pickup_point (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    option_pickups_id UUID NOT NULL REFERENCES option_pickups(id) ON DELETE CASCADE,
    pickup_point_id UUID NOT NULL REFERENCES pickup_point(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================================
-- UNIT RELATED TABLES
-- ============================================================================

-- Unit Content
CREATE TABLE unit_content (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT,
    title_plural TEXT,
    subtitle TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Unit
CREATE TABLE unit (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    option_id UUID NOT NULL REFERENCES option(id) ON DELETE CASCADE,
    internal_name TEXT NOT NULL,
    reference TEXT NOT NULL,
    type unit_type NOT NULL,
    min_age INTEGER NOT NULL,
    max_age INTEGER NOT NULL,
    id_required BOOLEAN NOT NULL DEFAULT false,
    min_quantity INTEGER,
    max_quantity INTEGER,
    pax_count INTEGER NOT NULL,
    unit_content_id UUID REFERENCES unit_content(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Unit Accompanied By
CREATE TABLE unit_accompanied_by (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    unit_id UUID NOT NULL REFERENCES unit(id) ON DELETE CASCADE,
    accompanied_by_unit_type unit_type NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Unit Contact Field
CREATE TABLE unit_contact_field (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    unit_id UUID NOT NULL REFERENCES unit(id) ON DELETE CASCADE,
    field contact_field NOT NULL,
    is_required BOOLEAN NOT NULL DEFAULT false,
    is_visible BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================================
-- OFFER RELATED TABLES
-- ============================================================================

-- Offer
CREATE TABLE offer (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_id UUID NOT NULL REFERENCES product(id) ON DELETE CASCADE,
    option_id UUID REFERENCES option(id) ON DELETE SET NULL,
    title TEXT NOT NULL,
    code TEXT NOT NULL,
    description TEXT,
    net_discount net_discount,
    usable BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Offer Restrictions
CREATE TABLE offer_restrictions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    offer_id UUID NOT NULL REFERENCES offer(id) ON DELETE CASCADE,
    min_units INTEGER,
    max_units INTEGER,
    min_total INTEGER,
    max_total INTEGER,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Offer Restrictions Unit
CREATE TABLE offer_restrictions_unit (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    offer_restrictions_id UUID NOT NULL REFERENCES offer_restrictions(id) ON DELETE CASCADE,
    unit_id UUID NOT NULL REFERENCES unit(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Tax
CREATE TABLE tax (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    type TEXT NOT NULL,
    description TEXT NOT NULL,
    amount INTEGER NOT NULL,
    percentage DOUBLE PRECISION,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Offer Discount
CREATE TABLE offer_discount (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    offer_id UUID NOT NULL REFERENCES offer(id) ON DELETE CASCADE,
    net INTEGER,
    retail INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Offer Discount Tax
CREATE TABLE offer_discount_tax (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    offer_discount_id UUID NOT NULL REFERENCES offer_discount(id) ON DELETE CASCADE,
    tax_id UUID NOT NULL REFERENCES tax(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Offer Combination
CREATE TABLE offer_combination (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_id UUID NOT NULL REFERENCES product(id) ON DELETE CASCADE,
    option_id UUID NOT NULL REFERENCES option(id) ON DELETE CASCADE,
    offer_code TEXT NOT NULL,
    offer_title TEXT NOT NULL,
    short_description TEXT,
    pricing_id UUID NOT NULL,
    booking_id UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Offer Combination Unit
CREATE TABLE offer_combination_unit (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    offer_combination_id UUID NOT NULL REFERENCES offer_combination(id) ON DELETE CASCADE,
    unit_id UUID NOT NULL REFERENCES unit(id) ON DELETE CASCADE,
    quantity INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================================
-- AVAILABILITY RELATED TABLES
-- ============================================================================

-- Tour Group
CREATE TABLE tour_group (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    internal_name TEXT NOT NULL,
    title TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Notice
CREATE TABLE notice (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT,
    short_description TEXT NOT NULL,
    cover_image_url TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Availability
CREATE TABLE availability (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    option_id UUID NOT NULL REFERENCES option(id) ON DELETE CASCADE,
    local_date_time_start TIMESTAMPTZ NOT NULL,
    local_date_time_end TIMESTAMPTZ NOT NULL,
    all_day BOOLEAN NOT NULL DEFAULT false,
    available BOOLEAN NOT NULL DEFAULT true,
    status availability_status NOT NULL,
    vacancies INTEGER,
    capacity INTEGER,
    max_units INTEGER,
    utc_cutoff_at TIMESTAMPTZ NOT NULL,
    meeting_point_id UUID REFERENCES meeting_point(id) ON DELETE SET NULL,
    meeting_local_date_time TIMESTAMPTZ,
    tour_group_id UUID REFERENCES tour_group(id) ON DELETE SET NULL,
    pickup_available BOOLEAN,
    pickup_required BOOLEAN,
    offer_code TEXT,
    offer_title TEXT,
    offer_id UUID REFERENCES offer(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Opening Hours
CREATE TABLE opening_hours (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    availability_id UUID NOT NULL REFERENCES availability(id) ON DELETE CASCADE,
    from_time TEXT NOT NULL,
    to_time TEXT NOT NULL,
    frequency TEXT,
    frequency_amount INTEGER,
    frequency_unit TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Availability Notice
CREATE TABLE availability_notice (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    availability_id UUID NOT NULL REFERENCES availability(id) ON DELETE CASCADE,
    notice_id UUID NOT NULL REFERENCES notice(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Availability Pickup Point
CREATE TABLE availability_pickup_point (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    availability_id UUID NOT NULL REFERENCES availability(id) ON DELETE CASCADE,
    pickup_point_id UUID NOT NULL REFERENCES pickup_point(id) ON DELETE CASCADE,
    local_date_time TEXT NOT NULL,
    local_date_time_to TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Availability Offer
CREATE TABLE availability_offer (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    availability_id UUID NOT NULL REFERENCES availability(id) ON DELETE CASCADE,
    offer_id UUID NOT NULL REFERENCES offer(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Resource Group
CREATE TABLE resource_group (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    split BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Resource
CREATE TABLE resource (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    resource_group_id UUID NOT NULL REFERENCES resource_group(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    seating BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Seat
CREATE TABLE seat (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    resource_id UUID NOT NULL REFERENCES resource(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    column_num INTEGER NOT NULL,
    row_num INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Availability Resource
CREATE TABLE availability_resource (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    availability_id UUID NOT NULL REFERENCES availability(id) ON DELETE CASCADE,
    resource_id UUID NOT NULL REFERENCES resource(id) ON DELETE CASCADE,
    quantity INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================================
-- BOOKING RELATED TABLES
-- ============================================================================

-- Contact
CREATE TABLE contact (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name TEXT,
    first_name TEXT,
    last_name TEXT,
    email_address TEXT,
    phone_number TEXT,
    postal_code TEXT,
    country TEXT,
    notes TEXT,
    allow_marketing BOOLEAN,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Contact Locale
CREATE TABLE contact_locale (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    contact_id UUID NOT NULL REFERENCES contact(id) ON DELETE CASCADE,
    locale TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Order
CREATE TABLE "order" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    test_mode BOOLEAN NOT NULL DEFAULT false,
    supplier_reference TEXT NOT NULL,
    settlement_method TEXT NOT NULL,
    status order_status NOT NULL,
    utc_expires_at TIMESTAMPTZ,
    utc_confirmed_at TIMESTAMPTZ,
    cancellable BOOLEAN NOT NULL DEFAULT false,
    contact_id UUID NOT NULL REFERENCES contact(id) ON DELETE RESTRICT,
    terms_accepted BOOLEAN,
    return_url TEXT,
    confirmable BOOLEAN,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Booking
CREATE TABLE booking (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    uuid UUID NOT NULL UNIQUE,
    test_mode BOOLEAN NOT NULL DEFAULT false,
    reseller_reference TEXT,
    supplier_reference TEXT,
    status booking_status NOT NULL,
    utc_created_at TIMESTAMPTZ NOT NULL,
    utc_updated_at TIMESTAMPTZ,
    utc_expires_at TIMESTAMPTZ,
    utc_redeemed_at TIMESTAMPTZ,
    utc_confirmed_at TIMESTAMPTZ,
    product_id UUID NOT NULL REFERENCES product(id) ON DELETE RESTRICT,
    option_id UUID NOT NULL REFERENCES option(id) ON DELETE RESTRICT,
    cancellable BOOLEAN NOT NULL DEFAULT false,
    freesale BOOLEAN NOT NULL DEFAULT false,
    availability_id UUID REFERENCES availability(id) ON DELETE SET NULL,
    contact_id UUID NOT NULL REFERENCES contact(id) ON DELETE RESTRICT,
    notes TEXT,
    meeting_point_id UUID REFERENCES meeting_point(id) ON DELETE SET NULL,
    meeting_local_date_time TIMESTAMPTZ,
    duration TEXT,
    duration_amount TEXT,
    duration_unit TEXT,
    terms_accepted BOOLEAN,
    pickup_requested BOOLEAN,
    pickup_point_id UUID REFERENCES pickup_point(id) ON DELETE SET NULL,
    pickup_hotel TEXT,
    pickup_hotel_room TEXT,
    order_id TEXT,
    order_reference TEXT,
    primary_booking BOOLEAN,
    offer_code TEXT,
    offer_title TEXT,
    offer_is_combination BOOLEAN,
    offer_id UUID REFERENCES offer(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Add foreign key for question_answer.booking_id
ALTER TABLE question_answer ADD CONSTRAINT fk_question_answer_booking 
    FOREIGN KEY (booking_id) REFERENCES booking(id) ON DELETE CASCADE;

-- Add foreign key for offer_combination.booking_id
ALTER TABLE offer_combination ADD CONSTRAINT fk_offer_combination_booking 
    FOREIGN KEY (booking_id) REFERENCES booking(id) ON DELETE SET NULL;

-- Order Booking
CREATE TABLE order_booking (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL REFERENCES "order"(id) ON DELETE CASCADE,
    booking_id UUID NOT NULL REFERENCES booking(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Cancellation
CREATE TABLE cancellation (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id UUID NOT NULL REFERENCES booking(id) ON DELETE CASCADE,
    refund TEXT NOT NULL,
    reason TEXT,
    utc_cancelled_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Booking Contact
CREATE TABLE booking_contact (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id UUID NOT NULL REFERENCES booking(id) ON DELETE CASCADE,
    contact_id UUID NOT NULL REFERENCES contact(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Ticket
CREATE TABLE ticket (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id UUID NOT NULL REFERENCES booking(id) ON DELETE CASCADE,
    redemption_method redemption_method NOT NULL,
    utc_redeemed_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Delivery Option
CREATE TABLE delivery_option (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ticket_id UUID NOT NULL REFERENCES ticket(id) ON DELETE CASCADE,
    delivery_format delivery_format NOT NULL,
    delivery_value TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Booking Delivery Method
CREATE TABLE booking_delivery_method (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id UUID NOT NULL REFERENCES booking(id) ON DELETE CASCADE,
    delivery_method delivery_method NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Booking Question Answer
CREATE TABLE booking_question_answer (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id UUID NOT NULL REFERENCES booking(id) ON DELETE CASCADE,
    question_answer_id UUID NOT NULL REFERENCES question_answer(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Booking Notice
CREATE TABLE booking_notice (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id UUID NOT NULL REFERENCES booking(id) ON DELETE CASCADE,
    notice_id UUID NOT NULL REFERENCES notice(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Unit Item
CREATE TABLE unit_item (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    uuid UUID NOT NULL UNIQUE,
    booking_id UUID NOT NULL REFERENCES booking(id) ON DELETE CASCADE,
    unit_id UUID NOT NULL REFERENCES unit(id) ON DELETE RESTRICT,
    reseller_reference TEXT,
    supplier_reference TEXT,
    status booking_status NOT NULL,
    utc_redeemed_at TIMESTAMPTZ,
    contact_id UUID REFERENCES contact(id) ON DELETE SET NULL,
    ticket_id UUID REFERENCES ticket(id) ON DELETE SET NULL,
    pricing_id UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Resource Allocation
CREATE TABLE resource_allocation (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id UUID NOT NULL REFERENCES booking(id) ON DELETE CASCADE,
    resource_group_id UUID NOT NULL REFERENCES resource_group(id) ON DELETE CASCADE,
    resource_id UUID NOT NULL REFERENCES resource(id) ON DELETE CASCADE,
    pax_count INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Resource Allocation Seat
CREATE TABLE resource_allocation_seat (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    resource_allocation_id UUID NOT NULL REFERENCES resource_allocation(id) ON DELETE CASCADE,
    seat_id UUID NOT NULL REFERENCES seat(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================================
-- PRICING RELATED TABLES
-- ============================================================================

-- Pricing
CREATE TABLE pricing (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    option_id UUID REFERENCES option(id) ON DELETE SET NULL,
    unit_id UUID REFERENCES unit(id) ON DELETE SET NULL,
    availability_id UUID REFERENCES availability(id) ON DELETE SET NULL,
    booking_id UUID REFERENCES booking(id) ON DELETE SET NULL,
    unit_item_id UUID REFERENCES unit_item(id) ON DELETE SET NULL,
    order_id UUID REFERENCES "order"(id) ON DELETE SET NULL,
    pricing_type TEXT NOT NULL,
    unit_type unit_type,
    original INTEGER NOT NULL,
    retail INTEGER NOT NULL,
    net INTEGER,
    currency TEXT NOT NULL,
    currency_precision INTEGER NOT NULL,
    offer_discount_id UUID REFERENCES offer_discount(id) ON DELETE SET NULL,
    extra_id UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Add foreign key for unit_item.pricing_id
ALTER TABLE unit_item ADD CONSTRAINT fk_unit_item_pricing 
    FOREIGN KEY (pricing_id) REFERENCES pricing(id) ON DELETE SET NULL;

-- Add foreign key for offer_combination.pricing_id
ALTER TABLE offer_combination ADD CONSTRAINT fk_offer_combination_pricing 
    FOREIGN KEY (pricing_id) REFERENCES pricing(id) ON DELETE RESTRICT;

-- Pricing Tax
CREATE TABLE pricing_tax (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    pricing_id UUID NOT NULL REFERENCES pricing(id) ON DELETE CASCADE,
    tax_id UUID NOT NULL REFERENCES tax(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Extra Pricing
CREATE TABLE extra_pricing (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    unit_id UUID NOT NULL REFERENCES unit(id) ON DELETE CASCADE,
    pricing_id UUID NOT NULL REFERENCES pricing(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================================
-- MAPPING RELATED TABLES
-- ============================================================================

-- Mapping
CREATE TABLE mapping (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    reseller_reference TEXT NOT NULL,
    reseller_status reseller_status NOT NULL,
    title TEXT NOT NULL,
    url TEXT NOT NULL,
    webhook_url TEXT,
    option_required BOOLEAN NOT NULL DEFAULT false,
    unit_required BOOLEAN NOT NULL DEFAULT false,
    product_id UUID REFERENCES product(id) ON DELETE SET NULL,
    option_id UUID REFERENCES option(id) ON DELETE SET NULL,
    unit_id UUID REFERENCES unit(id) ON DELETE SET NULL,
    connected BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- ============================================================================
-- INDEXES FOR PERFORMANCE
-- ============================================================================

-- Product indexes
CREATE INDEX idx_product_internal_name ON product(internal_name);
CREATE INDEX idx_product_reference ON product(reference);
CREATE INDEX idx_product_deleted_at ON product(deleted_at);

-- Option indexes
CREATE INDEX idx_option_product_id ON option(product_id);
CREATE INDEX idx_option_is_default ON option(is_default);
CREATE INDEX idx_option_deleted_at ON option(deleted_at);

-- Unit indexes
CREATE INDEX idx_unit_option_id ON unit(option_id);
CREATE INDEX idx_unit_type ON unit(type);
CREATE INDEX idx_unit_deleted_at ON unit(deleted_at);

-- Availability indexes
CREATE INDEX idx_availability_option_id ON availability(option_id);
CREATE INDEX idx_availability_status ON availability(status);
CREATE INDEX idx_availability_start_time ON availability(local_date_time_start);
CREATE INDEX idx_availability_end_time ON availability(local_date_time_end);
CREATE INDEX idx_availability_deleted_at ON availability(deleted_at);

-- Booking indexes
CREATE INDEX idx_booking_uuid ON booking(uuid);
CREATE INDEX idx_booking_status ON booking(status);
CREATE INDEX idx_booking_product_id ON booking(product_id);
CREATE INDEX idx_booking_option_id ON booking(option_id);
CREATE INDEX idx_booking_availability_id ON booking(availability_id);
CREATE INDEX idx_booking_contact_id ON booking(contact_id);
CREATE INDEX idx_booking_created_at ON booking(created_at);
CREATE INDEX idx_booking_deleted_at ON booking(deleted_at);

-- Order indexes
CREATE INDEX idx_order_status ON "order"(status);
CREATE INDEX idx_order_contact_id ON "order"(contact_id);
CREATE INDEX idx_order_deleted_at ON "order"(deleted_at);

-- Contact indexes
CREATE INDEX idx_contact_email ON contact(email_address);
CREATE INDEX idx_contact_deleted_at ON contact(deleted_at);

-- Offer indexes
CREATE INDEX idx_offer_product_id ON offer(product_id);
CREATE INDEX idx_offer_option_id ON offer(option_id);
CREATE INDEX idx_offer_code ON offer(code);
CREATE INDEX idx_offer_deleted_at ON offer(deleted_at);

-- Pricing indexes
CREATE INDEX idx_pricing_option_id ON pricing(option_id);
CREATE INDEX idx_pricing_unit_id ON pricing(unit_id);
CREATE INDEX idx_pricing_availability_id ON pricing(availability_id);
CREATE INDEX idx_pricing_booking_id ON pricing(booking_id);
CREATE INDEX idx_pricing_deleted_at ON pricing(deleted_at);

-- Mapping indexes
CREATE INDEX idx_mapping_product_id ON mapping(product_id);
CREATE INDEX idx_mapping_option_id ON mapping(option_id);
CREATE INDEX idx_mapping_unit_id ON mapping(unit_id);
CREATE INDEX idx_mapping_reseller_reference ON mapping(reseller_reference);
CREATE INDEX idx_mapping_deleted_at ON mapping(deleted_at);

