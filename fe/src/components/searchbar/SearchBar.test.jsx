import { describe, it, expect } from 'vitest'
import { render, screen } from '@testing-library/react'
import SearchBar from './SearchBar'

describe('SearchBar', () => {
  it('should render search input', () => {
    render(<SearchBar />)
    expect(screen.getByPlaceholderText(/search/i)).toBeInTheDocument()
  })

  it('should have correct input attributes', () => {
    render(<SearchBar />)
    const input = screen.getByPlaceholderText(/search/i)
    expect(input).toHaveAttribute('type', 'text')
    expect(input).toHaveClass('bg-[#1f2029]')
  })

  it('should render with correct container classes', () => {
    render(<SearchBar />)
    const container = screen.getByPlaceholderText(/search/i).parentElement
    expect(container).toHaveClass('relative')
  })
})